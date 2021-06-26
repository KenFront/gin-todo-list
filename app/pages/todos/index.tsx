import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";
import { SignOutButton } from "@/auth/SignOutButton";
import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";

export const getServerSideProps = CheckPageWithAuth;

const TodolistPage = () => {
  return (
    <FullPage>
      <Header title="Todo list" rightArea={<SignOutButton />} />
      <Responsive>Hello world</Responsive>
    </FullPage>
  );
};

export default TodolistPage;
