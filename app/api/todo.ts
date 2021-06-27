import { Request, PostOptions } from "@/lib/request";
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
  const res = await Request({
    path: "/api/todos",
    options: PostOptions(val),
  });

  return res as ApiOnSuccess<Todo>;
};

export const getTodos = async () => {
  const res = await Request({
    path: "/api/todos",
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
  const res = await Request({
    path: `http://localhost/api/todos/${id}`,
    options: {
      headers: {
        cookie,
      },
    },
  });

  return res as ApiOnSuccess<Todo>;
};
