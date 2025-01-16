import { SHA256, enc } from 'crypto-js';

export const hashPassword = (password: string): string => {
  // Хеширование пароля с использованием SHA-256
  const hash = SHA256(password).toString(enc.Hex);
  return hash;
};

export const comparePassword = (providedPass: string, storedPass: string): boolean => {
  // Сравнение хешей
  const providedHash = SHA256(providedPass).toString(enc.Hex);
  return providedHash === storedPass;
};
