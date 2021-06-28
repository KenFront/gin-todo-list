import { useEffect, FC } from "react";
import { GetServerSideProps } from "next";
import { Box, Button, Stack } from "@chakra-ui/react";
import { Formik, Form, useFormikContext } from "formik";
import { format } from "date-fns";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";

import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";
import { EditableText } from "@/lib/component/formik/EditableText";
import { Text } from "@/lib/component/formik/Text";
import { Switch } from "@/lib/component/formik/Switch";
import { SignOutButton } from "@/auth/SignOutButton";
import { ErrorPage } from "@/lib/component/ErrorPage";

import { useAsync } from "@/lib/hook/useAsync";
import { useAppToast } from "@/lib/hook/useAppToast";

import { getTodoByIdOnServerSide, patchTodoById, Todo } from "@/api/todo";
import { GetErrorHandler, UnPromisify } from "@/lib/request";

import { isNotEmpty } from "@/validator/isNotEmpty";

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const checkAuth = await CheckPageWithAuth(ctx);
  if (checkAuth.hasOwnProperty("redirect")) {
    return checkAuth;
  }

  const id = ctx.params?.id;

  if (typeof id === "string") {
    try {
      const res = await getTodoByIdOnServerSide({
        id,
        cookie: ctx.req.headers.cookie ?? "",
      });
      return {
        props: {
          res,
        },
      };
    } catch (e) {
      return {
        props: {
          res: e,
        },
      };
    }
  }
  return {
    props: {},
  };
};

const UpdateInitValues: FC<{ todo: Todo }> = ({ todo }) => {
  const { resetForm } = useFormikContext();

  useEffect(() => {
    if (todo) {
      resetForm({
        values: {
          title: todo.title,
          description: todo.description,
          createdAt: format(new Date(todo.createdAt), "yyyy-MM-dd HH:mm:ss z"),
          updatedAt: format(new Date(todo.updatedAt), "yyyy-MM-dd HH:mm:ss z"),
          isCompleted: todo.status === "completed",
        },
      });
    }
  }, [resetForm, todo]);
  return null;
};

const TodoDetailPage = ({
  res,
}: {
  res: UnPromisify<ReturnType<typeof getTodoByIdOnServerSide>>;
}) => {
  const { status: apiStatus, result, error, execute } = useAsync(patchTodoById);
  const { toastSuccess, toastError } = useAppToast();

  useEffect(() => {
    if (apiStatus === "success" && result) {
      toastSuccess({
        title: "Success",
        description: "Update todo successfully",
      });
    }
  }, [apiStatus, result, toastSuccess]);

  useEffect(() => {
    if (apiStatus === "error" && error) {
      const e = GetErrorHandler(error);
      toastError({
        title: "Error",
        description: e,
      });
    }
  }, [apiStatus, error, toastError]);

  const e = GetErrorHandler(res);
  if (e) {
    return (
      <ErrorPage title="Todo Detail" rightArea={<SignOutButton />} msg={e} />
    );
  }

  const { id, title, description, createdAt, updatedAt, status } = res.data;

  return (
    <FullPage>
      <Header title="Todo Add" rightArea={<SignOutButton />} />
      <Responsive align="center" justify="center">
        <Box w="480px" p={4}>
          <Formik
            initialValues={{
              title,
              description,
              createdAt: format(new Date(createdAt), "yyyy-MM-dd HH:mm:ss z"),
              updatedAt: format(new Date(updatedAt), "yyyy-MM-dd HH:mm:ss z"),
              isCompleted: status === "completed",
            }}
            onSubmit={(values) => {
              execute({
                id,
                title: values.title,
                description: values.description,
                status: values.isCompleted ? "completed" : "idle",
              });
            }}
          >
            <Form>
              <EditableText
                name="title"
                label="Title"
                validate={isNotEmpty("Title")}
                placeholder="title"
              />
              <EditableText
                name="description"
                label="Description"
                validate={isNotEmpty("Description")}
                placeholder="description"
              />
              <Text name="createdAt" label="CreatedAt" />
              <Text name="updatedAt" label="UpdatedAt" />
              <Switch name="isCompleted" label="Status" />
              <Stack mt={4} direction="row" spacing={4} align="center">
                <Button
                  colorScheme="teal"
                  isLoading={apiStatus === "loading"}
                  type="submit"
                >
                  Submit
                </Button>
              </Stack>
              {apiStatus === "success" && result && (
                <UpdateInitValues todo={result.data} />
              )}
            </Form>
          </Formik>
        </Box>
      </Responsive>
    </FullPage>
  );
};

export default TodoDetailPage;
