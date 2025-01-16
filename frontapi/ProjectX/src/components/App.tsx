import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import { useAppState } from "./ctx/AppState";
import { hashPassword } from "../utils/Hash";

const App: React.FC = () => {
  const [login, setLogin] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [loginError, setLoginError] = useState<string | null>(null); // Состояние для ошибки
  const { setError, setMessage } = useAppState();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent): Promise<void> => {
    e.preventDefault();

    try {
      // Хеширование пароля
      const hashedPassword = await hashPassword(password);
      console.log("Hashed Password:", hashedPassword); // Выводим хеш пароля в консоль

      // Отправка запроса на сервер с логином и паролем
      const response = await axios.post("http://localhost:8080/front/v1/auth", {
        login,
        password: hashedPassword,
      });

      // Обработка ответа от сервера
      if (response.status === 200) {
        setMessage("Login successful!");
        navigate("/login-success");
      } else if (response.status === 403) {
        setMessage("Invalid login or password");
        setLoginError("Invalid login or password");
        setError("403");
      } else if (response.status === 404) {
        setMessage("Invalid login or password");
        setLoginError("Invalid login or password");
        setError("404");
      } else {
        setMessage("An error occurred, please try again.");
        setError("500");
        navigate("/error");
      }
    } catch (err) {
      console.error("Error during login:", err); // Логирование ошибок
      setMessage("Не корректные логин или пароль");
      setError("500");
      navigate("/error");
    }
  };

  return (
    <div className="container">
      <h1>Добро пожаловать в ProjectX</h1>
      <div className="form-container">
        <h2>Вход</h2>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Логин:</label>
            <input
              type="text"
              value={login}
              onChange={(e) => setLogin(e.target.value)}
            />
          </div>
          <div className="form-group">
            <label>Пароль:</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          {/* Отображаем ошибку, если она есть */}
          {loginError && <div className="error-message">{loginError}</div>}

          <div className="form-actions">
            <button type="submit">Вход</button>
            <button type="button" onClick={() => navigate("/register")}>
              Регистрация
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default App;
