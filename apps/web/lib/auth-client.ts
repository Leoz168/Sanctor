"use client";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

type AuthResponse = {
  token: string;
  refreshToken?: string;
  expiresAt?: string;
};

type LoginPayload = {
  email: string;
  password: string;
};

type RegisterPayload = {
  email: string;
  password: string;
  username: string;
  firstName: string;
  lastName: string;
};

function readErrorMessage(data: unknown, fallback: string) {
  if (typeof data === "object" && data !== null && "error" in data && typeof data.error === "string") {
    return data.error;
  }

  return fallback;
}

async function sendAuthRequest<TPayload>(path: string, payload: TPayload, fallbackError: string) {
  const response = await fetch(`${apiBase}${path}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });

  const data = (await response.json().catch(() => null)) as AuthResponse | { error?: string } | null;
  if (!response.ok || !data || typeof (data as AuthResponse).token !== "string") {
    throw new Error(readErrorMessage(data, fallbackError));
  }

  return data as AuthResponse;
}

export function persistAuthToken(token: string) {
  localStorage.setItem("authToken", token);
  localStorage.setItem("token", token);
}

export async function loginWithPassword(payload: LoginPayload) {
  return sendAuthRequest("/api/auth/login", payload, "Login failed.");
}

export async function registerWithPassword(payload: RegisterPayload) {
  return sendAuthRequest("/api/auth/register", payload, "Registration failed.");
}

export function splitFullName(fullName: string) {
  const parts = fullName.trim().split(/\s+/).filter(Boolean);
  if (parts.length === 0) {
    return { firstName: "", lastName: "" };
  }

  return {
    firstName: parts[0],
    lastName: parts.slice(1).join(" "),
  };
}

export function buildUsername(fullName: string, email: string) {
  const emailPrefix = email.trim().split("@")[0]?.toLowerCase() ?? "";
  if (emailPrefix) {
    return emailPrefix.replace(/[^a-z0-9._-]/g, "");
  }

  return fullName.trim().toLowerCase().replace(/[^a-z0-9]+/g, "");
}
