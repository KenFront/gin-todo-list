import { FC } from "react";
import { Flex } from "@chakra-ui/react";

export const FullPage: FC = (props) => (
  <Flex w="100vw" h="100vh" flexWrap="nowrap" flexDirection="column">
    {props.children}
  </Flex>
);
