import { FC } from "react";
import { Flex, FlexProps } from "@chakra-ui/react";

export const Responsive: FC<FlexProps> = ({ children, ...flexProps }) => (
  <Flex
    flexGrow={1}
    flexShrink={1}
    w="100%"
    flexBasis="auto"
    flexWrap="wrap"
    overflowX="hidden"
    overflowY="auto"
    {...flexProps}
  >
    {children}
  </Flex>
);
