import { Heading } from "@chakra-ui/react";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";

import { FullPage } from "@/lib/component/FullPage";
import { Responsive } from "@/lib/component/Responsive";
import { Header } from "@/lib/component/Header";

export const getServerSideProps = CheckPageWithAuth;

const TodoDetailPage = () => {
  return (
    <FullPage>
      <Header title="Todo Detail" />
      <Responsive align="center" justify="center">
        <Heading size="lg">Hello World</Heading>
      </Responsive>
    </FullPage>
  );
};

export default TodoDetailPage;
