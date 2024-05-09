import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import toast from "react-hot-toast";
import { IoMdAdd } from "react-icons/io";

const TodoForm = () => {
  const [newTodo, setNewTodo] = useState("");

  const queryClient = useQueryClient();

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (event: React.FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      try {
        const url = `${import.meta.env.VITE_BASE_URL}/todos`;
        const res = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ body: newTodo }),
        });

        const data = await res.json();

        if (!res.ok) {
          toast.error("Something went wrong while adding new Todo");
          throw new Error(
            data.error || "Something went wrong while adding new Todo"
          );
        }
        setNewTodo("");

        toast.success("Todo added successfuly");
      } catch (e) {
        console.log(e);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: () => {
      toast.error("Add the todo again!");
    },
  });

  return (
    <form onSubmit={createTodo}>
      <Flex gap={2}>
        <Input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
        />
        <Button
          mx={2}
          type="submit"
          _active={{
            transform: "scale(.97)",
          }}
        >
          {isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
        </Button>
      </Flex>
    </form>
  );
};
export default TodoForm;
