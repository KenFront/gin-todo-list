import { MouseEventHandler } from "react";
import {
  Wrap,
  WrapItem,
  LinkBox,
  LinkOverlay,
  Heading,
} from "@chakra-ui/react";

import { CheckPageWithAuth } from "@/auth/CheckPageWithAuth";
import { SignOutButton } from "@/auth/SignOutButton";

import { FullPage } from "@/lib/component/FullPage";
import { Header } from "@/lib/component/Header";
import { Responsive } from "@/lib/component/Responsive";

import { useChangePage } from "@/lib/hook/useChangePage";
import { ROUTE } from "@/route";

export const getServerSideProps = CheckPageWithAuth;

const IndexPage = () => {
  const { changePath } = useChangePage();

  const toTodoListPage: MouseEventHandler<HTMLAnchorElement> = (e) => {
    e.preventDefault();
    changePath(ROUTE.TODOS);
  };

  return (
    <FullPage>
      <Header title="Index" rightArea={<SignOutButton />} />
      <Responsive>
        <Wrap p={4} w="100%" spacing={4} justify="center">
          <WrapItem>
            <LinkBox
              as="article"
              maxW="sm"
              p="4"
              borderWidth="1px"
              rounded="md"
            >
              <Heading size="md" my="2">
                <LinkOverlay href={ROUTE.TODOS} onClick={toTodoListPage}>
                  TodoList
                </LinkOverlay>
              </Heading>
            </LinkBox>
          </WrapItem>
          <WrapItem>
            <LinkBox
              as="article"
              maxW="sm"
              p="4"
              borderWidth="1px"
              rounded="md"
            >
              <Heading size="md" my="2">
                <LinkOverlay>Profile</LinkOverlay>
              </Heading>
            </LinkBox>
          </WrapItem>
        </Wrap>
      </Responsive>
    </FullPage>
  );
};

export default IndexPage;
