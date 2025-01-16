import React, { createContext, useState, useContext } from "react";

// Тип для состояния контекста
type AppStateType = {
  error: string | null;
  message: string | null;
  setError: (error: string | null) => void;
  setMessage: (message: string | null) => void;
};

// Создание контекста
const AppStateContext = createContext<AppStateType | undefined>(undefined);

// Провайдер контекста
export const AppStateProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [error, setError] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);

  return (
    <AppStateContext.Provider value={{ error, message, setError, setMessage }}>
      {children}
    </AppStateContext.Provider>
  );
};

// Хук для использования контекста
export const useAppState = (): AppStateType => {
  const context = useContext(AppStateContext);
  if (!context) {
    throw new Error("useAppState must be used within an AppStateProvider");
  }
  return context;
};
