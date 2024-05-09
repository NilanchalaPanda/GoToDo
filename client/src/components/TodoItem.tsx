import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { Todo } from "./TodoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";
import { BsDashCircleDotted } from "react-icons/bs";
import { BASEURL } from "../App";

const TodoItem = ({ todo }: { todo: Todo }) => {
  const queryClient = useQueryClient();
  const { mutate: updateTodo, isPending: isUpdating } = useMutation({
    mutationKey: ["updateTodo"],

    mutationFn: async () => {
      if (todo.completed) toast.success("Already updated");
      try {
        const url = `${BASEURL}/todos/${todo._id}`;
        const res = await fetch(url, {
          method: "PATCH",
        });
        const data = await res.json();

        if (!res.ok) {
          toast.error("Something went wrong while updating");
          throw new Error(data.error || "Something went wrong while updating");
        }
      } catch (err) {
        console.log("UpdatingError -", err);
      }
    },

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  
  const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
    mutationKey: ["deleteTodo"],

    mutationFn: async () => {
      try {
        const url = `${BASEURL}/todos/${todo._id}`;
        const res = await fetch(url, {
          method: "DELETE",
        });
        const data = await res.json();

        if (!res.ok) {
          toast.error("Something went wrong while updating");
          throw new Error(data.error || "Something went wrong while deleting");
        }
      } catch (err) {
        console.log("DeletingError -", err);
      }
    },

    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  return (
    <Flex gap={2} alignItems={"center"}>
      <Flex
        flex={1}
        alignItems={"center"}
        border={"1px"}
        borderColor={"gray.600"}
        p={2}
        borderRadius={"lg"}
        justifyContent={"space-between"}
      >
        <Text
          color={todo.completed ? "green.200" : "yellow.100"}
          textDecoration={todo.completed ? "line-through" : "none"}
        >
          {todo.body}
        </Text>
        {todo.completed && (
          <Badge ml="1" colorScheme="green">
            Done
          </Badge>
        )}
        {!todo.completed && (
          <Badge ml="1" colorScheme="yellow">
            In Progress
          </Badge>
        )}
      </Flex>
      <Flex gap={2} alignItems={"center"}>
        <Box
          color={"green.500"}
          cursor={"pointer"}
          onClick={() => updateTodo()}
        >
          {!todo.completed && <BsDashCircleDotted size={20} />}
          {todo.completed && isUpdating && <Spinner size={"sm"} />}
          {todo.completed && !isUpdating && <FaCheckCircle size={20} />}
        </Box>
        <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTodo()}>
          {!isDeleting && <MdDelete size={25} />}
          {isDeleting && <Spinner size={"sm"} />}
        </Box>
      </Flex>
    </Flex>
  );
};
export default TodoItem;
