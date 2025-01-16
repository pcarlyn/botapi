import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css'; // Импорт CSS
import Root from './pages/Root';

const rootElement = document.getElementById('root');
if (rootElement) {
  createRoot(rootElement).render(
    <StrictMode>
      <Root />
    </StrictMode>
  );
} else {
  console.error('Element with id "root" not found');
}
