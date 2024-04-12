import React, { createContext, useState, useEffect } from 'react';

export const WebsocketContext = createContext({
  conn: null,
  setConn: () => {},
});

const WebsocketProvider = ({ children }) => {
  const [conn, setConn] = useState(null);

  useEffect(() => {
    const cleanup = () => {
      if (conn !== null) {
        conn.close();
      }
    };

    return cleanup;
  }, [conn]);

  return (
    <WebsocketContext.Provider value={{ conn, setConn }}>
      {children}
    </WebsocketContext.Provider>
  );
};

export { WebsocketProvider };
