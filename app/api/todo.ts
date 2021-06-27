import { ClientRequest, ServerRequest, PostOptions } from "@/lib/request";
import { ApiOnSuccess } from "@/lib/request/type";

export interface Todo {
  id: string;
  title: string;
  description: string;
  status: "idle" | "completed";
  createdAt: string;
  updatedAt: string;
  userId: string;
}

export const addTodo = async (val: { title: string; description: string }) => {
  const res = await ClientRequest({
    path: "/todos",
    options: PostOptions(val),
  });

  return res as ApiOnSuccess<Todo>;
};

export const getTodos = async () => {
  const res = await ClientRequest({
    path: "/todos",
  });

  return res as ApiOnSuccess<Todo[]>;
};

export const getTodoByIdOnServerSide = async ({
  id,
  cookie,
}: {
  id: string;
  cookie: string;
}) => {
  const res = await ServerRequest({
    path: `/todos/${id}`,
    options: {
      headers: {
        cookie,
      },
    },
  });

  return res as ApiOnSuccess<Todo>;
};
