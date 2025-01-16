import React from "react";
import { useAppState } from "./ctx/AppState";

const ErrorPage: React.FC = () => {
  const { error, message } = useAppState();

  return (
    <div className="container">
      <div className="form-container">
        <h1>Error {error || "X3"}</h1>
        <p>{message || "Что-то не так"}</p>
        <a href="/">Вернуться</a>
      </div>
    </div>
  );
};

export default ErrorPage;
