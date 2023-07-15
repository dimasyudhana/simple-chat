import { createContext, useState } from "react";

const AuthContext = createContext({});

export const AuthProvider = ({ children }) => {
    const [auth, setAuthentication] = useState({});

    return (
        <AuthContext.Provider value={{ auth, setAuthentication }}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContext;