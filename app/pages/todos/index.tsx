import { useEffect, useMemo, useState } from "react";
import { Table, Column } from "react-virtualized";
import { Checkbox, HStack, Button, Stack, Skeleton } from "@chakra-ui/react";
import { EditIcon, DeleteIcon } from "@chakra-ui/icons";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";
import { SignOutButton } from "@/auth/SignOutButton";
import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";
import { DeleteModal } from "@/lib/component/DeleteModal";

import { useAsync } from "@/lib/hook/useAsync";
import { useChangePage } from "@/lib/hook/useChangePage";
import { useAppToast } from "@/lib/hook/useAppToast";

import { getFormatedTime } from "@/lib/time/format";

import { GetErrorHandler } from "@/lib/request";
import { getTodos, deleteTodoById } from "@/api/todo";

import { ROUTE } from "@/route";

export const getServerSideProps = CheckPageWithAuth;

const TodolistPage = () => {
  const { execute, result } = useAsync(getTodos);
  const [deleteId, setDeleteId] = useState("");
  const { changePath } = useChangePage();
  const { toastSuccess, toastError } = useAppToast();

  const toTodoAddPage = () => {
    changePath({ path: ROUTE.TODO_ADD });
  };

  const toTodoDetailPage = (id: string) => {
    changePath({ path: ROUTE.TODO_DETAIL, param: { id } });
  };

  const colW = useMemo(() => [100, 200, 100, 250, 250, 80], []);
  const colSum = colW.reduce((sum, num) => sum + num);

  useEffect(() => {
    execute();
  }, [execute]);

  return (
    <FullPage>
      <Header title="Todo list" rightArea={<SignOutButton />} />
      <Responsive p={4} justify="center">
        <HStack w={colSum} spacing={4}>
          <Button
            leftIcon={<EditIcon />}
            colorScheme="facebook"
            variant="solid"
            onClick={toTodoAddPage}
          >
            Add
          </Button>
        </HStack>
        {result ? (
          <Table
            width={colSum}
            height={360}
            headerHeight={28}
            rowHeight={60}
            rowCount={result.data.length}
            rowGetter={({ index }) => result.data[index]}
            onRowClick={({ index }) => toTodoDetailPage(result.data[index].id)}
            gridStyle={{
              cursor: "pointer",
            }}
          >
            <Column width={colW[0]} label="Title" dataKey="title" />
            <Column width={colW[1]} label="Description" dataKey="description" />
            <Column
              width={colW[2]}
              label="Status"
              dataKey="status"
              cellRenderer={({ cellData }) => (
                <Checkbox isChecked={cellData === "completed"} isDisabled />
              )}
            />
            <Column
              width={colW[3]}
              label="CreatedAt"
              dataKey="createdAt"
              cellRenderer={({ cellData }) => (
                <span>{getFormatedTime(cellData)}</span>
              )}
            />
            <Column
              width={colW[4]}
              label="UpdatedAt"
              dataKey="updatedAt"
              cellRenderer={({ cellData }) => (
                <span>{getFormatedTime(cellData)}</span>
              )}
            />
            <Column
              width={colW[5]}
              label="Action"
              dataKey="id"
              cellRenderer={({ cellData }) => (
                <Button
                  colorScheme="red"
                  onClick={(e) => {
                    e.stopPropagation();
                    setDeleteId(cellData);
                  }}
                >
                  <DeleteIcon />
                </Button>
              )}
            />
          </Table>
        ) : (
          <Stack w={colSum}>
            {new Array(18).fill("").map((_val, i) => (
              <Skeleton key={i} height="20px" />
            ))}
          </Stack>
        )}
      </Responsive>
      <DeleteModal
        isOpen={!!deleteId}
        title="Delete todo"
        description={`Are you sure to delete ${
          result?.data.find((item) => item.id === deleteId)?.title || "it"
        }?`}
        confirmText="Ok"
        cancelText="cancel"
        onConfirm={async () => {
          try {
            const res = await deleteTodoById({ id: deleteId });
            toastSuccess({
              title: "Success",
              description: `Delete ${res.data.title} successfully`,
            });
            setDeleteId("");
          } catch (e) {
            const msg = GetErrorHandler(e);
            toastError({
              title: "Error",
              description: msg,
            });
          } finally {
            execute();
          }
        }}
        onCancel={async () => {
          setDeleteId("");
        }}
      />
    </FullPage>
  );
};

export default TodolistPage;
