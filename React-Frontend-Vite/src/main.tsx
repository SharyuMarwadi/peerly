import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { ThemeProvider } from '@emotion/react'
import theme from './theme.tsx'
import { GoogleOAuthProvider } from '@react-oauth/google';

ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <ThemeProvider theme={theme}>
            <GoogleOAuthProvider clientId="859041929183-p1mr0a0qebijd8ii4nprkb5j8ki506op.apps.googleusercontent.com">
                <App />
            </GoogleOAuthProvider>
        </ThemeProvider>
    </React.StrictMode >,
)
