import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { ChakraProvider } from "@chakra-ui/react";
import { Toaster } from "react-hot-toast";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import theme from "./chakra/theme.ts";

const queryclient = new QueryClient();
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryclient}>
      <ChakraProvider theme={theme}>
        <App />
        <Toaster position="bottom-center" />
      </ChakraProvider>
    </QueryClientProvider>
  </React.StrictMode>
);
