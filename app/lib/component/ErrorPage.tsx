import { FC, ReactNode } from "react";
import { Heading } from "@chakra-ui/react";

import { FullPage } from "./FullPage";
import { Responsive } from "./Responsive";
import { Header } from "./Header";

export const ErrorPage: FC<{
  msg: string;
  title: string;
  rightArea?: ReactNode;
}> = ({ msg, title, rightArea }) => {
  return (
    <FullPage>
      <Header title={title} rightArea={rightArea} />
      <Responsive align="center" justify="center">
        <Heading size="lg">{msg}</Heading>
      </Responsive>
    </FullPage>
  );
};
