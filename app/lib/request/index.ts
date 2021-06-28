export const DEFAULT_REQUEST_CONFIG: RequestInit = {
  cache: "no-cache",
  credentials: "same-origin",
  headers: {
    "content-type": "application/json",
    "X-Service": "gin_server",
  },
  redirect: "follow",
  referrer: "no-referrer",
};
Object.freeze(DEFAULT_REQUEST_CONFIG);

export type UnPromisify<T> = T extends Promise<infer U> ? U : T;

export const CLIENT_REQUEST_PATH_PREFIX =
  process.env.NODE_ENV === "production" ? "" : "/api";
export const SERVER_REQUEST_PATH_PREFIX = `http://${process.env.DOMAIN}`;

const request = async ({
  path,
  options,
}: {
  path: string;
  options?: RequestInit;
}) => {
  const response = await fetch(path, options);

  let jsonRes: unknown;

  try {
    jsonRes = await response.json();
  } catch (e) {
    throw {
      error: "Something wrong",
    };
  }

  if (response.ok) {
    return jsonRes;
  } else {
    throw jsonRes;
  }
};

export const ClientRequest = ({
  path,
  options,
}: Parameters<typeof request>[0]) => {
  return request({
    path: `${CLIENT_REQUEST_PATH_PREFIX}${path}`,
    options: { ...DEFAULT_REQUEST_CONFIG, ...options },
  });
};

export const ServerRequest = ({
  path,
  options,
}: Parameters<typeof request>[0]) => {
  return request({
    path: `${SERVER_REQUEST_PATH_PREFIX}${path}`,
    options: {
      headers: {
        ...DEFAULT_REQUEST_CONFIG.headers,
        ...options?.headers
      }
    },
  });
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const GetErrorHandler = (e: any): string => {
  if (e && e.error && typeof e.error === "string") {
    console.error(e);
    return e.error;
  } else {
    return "";
  }
};

export const PostOptions = (body?: unknown): RequestInit => ({
  method: "post",
  body: body ? JSON.stringify(body) : null,
});
export const PatchOptions = (body?: unknown): RequestInit => ({
  method: "PATCH",
  body: body ? JSON.stringify(body) : null,
});
