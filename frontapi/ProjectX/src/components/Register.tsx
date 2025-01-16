import React, { useState } from 'react';
import axios from 'axios';
import { hashPassword } from '../utils/Hash';

const Register: React.FC = () => {
  const [firstname, setFirstname] = useState(''); // имя
  const [secondname, setSecondname] = useState(''); // Фамилия
  const [userborn, setUserborn] = useState(''); // дата рождения
  const [usersex, setUsersex] = useState(''); // пол
  const [usersity, setUsersity] = useState(''); // город
  const [username, setUsername] = useState(''); // Логин
  const [password, setPassword] = useState(''); // Пароль
  const [captcha, setCaptcha] = useState(''); // капча
  const [confirmPassword, setConfirmPassword] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [successMessage, setSuccessMessage] = useState('');

  const isFirstnameValid = (firstname: string) => {
    const isRussianOnly = /^[а-яА-ЯёЁ\s]+$/;
    return isRussianOnly.test(firstname);
  };

  const isSecondnameValid = (secondname: string) => {
    const isRussianOnly = /^[а-яА-ЯёЁ\s]+$/;
    return isRussianOnly.test(secondname);
  };

  const isUserbornValid = (userborn: string) => {
    const datePattern = /^\d{2}-\d{2}-\d{4}$/; 
    return datePattern.test(userborn);
  };

  const isLoginValid = (username: string) => {
    const isEnglishOnly = /^[a-zA-Z0-9]+$/;
    return isEnglishOnly.test(username);
  };

  const isPasswordValid = (password: string) => {
    const minLength = /.{8,}/;
    const hasLetter = /[a-zA-Z]/;
    const hasDigit = /\d/;
    const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/;

    return (
      minLength.test(password) &&
      hasLetter.test(password) &&
      hasDigit.test(password) &&
      hasSpecialChar.test(password)
    );
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!isFirstnameValid(firstname)) {
      setErrorMessage('Имя должно содержать только русские буквы.');
      return;
    }

    if (!isSecondnameValid(secondname)) {
      setErrorMessage('Фамилия должна содержать только русские буквы.');
      return;
    }

    if (!isUserbornValid(userborn)) {
      setErrorMessage('Дата рождения должна быть в формате DD-MM-YYYY.');
      return;
    }

    // Валидация имени пользователя
    if (!isLoginValid(username)) {
      setErrorMessage('Логин должен содержать только английские буквы и цифры.');
      return;
    }

    // Валидация пароля
    if (!isPasswordValid(password)) {
      setErrorMessage('Пароль должен быть длиной не менее 8 символов и содержать буквы, цифры и специальные символы.');
      return;
    }

    // Проверка совпадения паролей
    if (password !== confirmPassword) {
      setErrorMessage('Пароли не совпадают!');
      return;
    }

    try {
      // Хеширование пароля
      const hashedPassword = await hashPassword(password);
      
      // Отправка данных на сервер
      await axios.post('/api/register', { username: username.trim(), password: hashedPassword });
      
      setSuccessMessage('Регистрация прошла успешно!');
      setErrorMessage('');
      
      // Сброс полей после успешной регистрации
      setFirstname('');
      setSecondname('');
      setUsername('');
      setPassword('');
      setConfirmPassword('');
      
    } catch (error) {
      console.error('Ошибка при регистрации:', error);
      
      if (axios.isAxiosError(error) && error.response) {
        setErrorMessage(error.response.data.message || 'Ошибка регистрации. Попробуйте еще раз.');
      } else {
        setErrorMessage('Ошибка сети или сервера.');
      }
    }
  };

  return (
    <div className="container">
      <h1>Регистрация</h1>
      <form className="form-container" onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Имя</label>
          <input
            type="text"
            value={firstname}
            onChange={(e) => setFirstname(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Фамилия</label>
          <input
            type="text"
            value={secondname}
            onChange={(e) => setSecondname(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Дата рождения</label>
          <input
            type="text"
            value={userborn}
            onChange={(e) => setUserborn(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Город</label>
          <input
            type="text"
            value={usersity}
            onChange={(e) => setUsersity(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Пол</label>
          <input
            type="text"
            value={usersex}
            onChange={(e) => setUsersex(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Логин</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Пароль</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Подтвердите пароль</label>
          <input
            type="password"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            required />
        </div>
        <div className="form-group">
          <label>Введите капчу</label>
          <input
            type="text"
            value={captcha}
            onChange={(e) => setCaptcha(e.target.value)}
            required />
        </div>

        {errorMessage && <p className="error-message">{errorMessage}</p>}
        {successMessage && <p className="success-message">{successMessage}</p>}
        <div className="form-actions">
          <button type="submit">Зарегистрироваться</button>
        </div>
      </form>
    </div>
  );
};

export default Register;
