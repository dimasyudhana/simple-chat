// App.js
import React from 'react';
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom';
import { AuthProvider } from './context/AuthProvider';
import { WebsocketProvider } from './context/WebsocketProvider';
import Login from './Login';
import Register from './Register';
import Homepage from './Homepage';

function App() {
    return (
        <BrowserRouter>
            <AuthProvider>
                <WebsocketProvider>
                    <main className='App'>
                        <Routes>
                            <Route path='/login' element={<Login />} />
                            <Route path='/signup' element={<Register />} />
                            <Route path='/homepage' element={<Homepage />} />
                            <Route path='/' element={<Navigate to='/login' replace />} />
                        </Routes>
                    </main>
                </WebsocketProvider>
            </AuthProvider>
        </BrowserRouter>
    );
}

export default App;
