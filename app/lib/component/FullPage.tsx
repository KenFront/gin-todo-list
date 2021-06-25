import { FC } from "react";
import { Flex } from "@chakra-ui/react";

export const FullPage: FC = (props) => (
  <Flex w="100vw" h="100vh" align="center" justify="center" flexWrap="wrap">
    {props.children}
  </Flex>
);
