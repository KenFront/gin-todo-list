import { format } from "date-fns";

export const getFormatedTime = (time: string) =>
  format(new Date(time), "yyyy-MM-dd HH:mm:ss z");
