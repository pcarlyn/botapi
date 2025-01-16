import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import App from "../components/App";
import LoginSuccess from "../components/LoginSuccess";
import Register from "../components/Register";
import ErrPage from "../components/ErrorPage";
import { AppStateProvider } from "../components/ctx/AppState";

const Root: React.FC = () => {
  return (
    <AppStateProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/login-success" element={<LoginSuccess />} />
          <Route path="/register" element={<Register />} />
          <Route path="/error" element={<ErrPage />} />
        </Routes>
      </BrowserRouter>
    </AppStateProvider>
  );
};

export default Root;
