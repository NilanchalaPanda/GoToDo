import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

// eslint-disable-next-line react-refresh/only-export-components
export const BASEURL =
  import.meta.env.MODE === "development"
    ? import.meta.env.VITE_BASE_URL
    : "/api";

function App() {
  return (
    <Stack h="100vh" mb="100px">
      <Navbar />
      <Container>
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  );
}

export default App;
