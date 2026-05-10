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

export type CurrentUser = {
  id: string;
  email: string;
  username: string;
  avatar?: string;
  bio?: string;
  gender?: string;
  age?: number | null;
  institutionId?: string | null;
  major?: string | null;
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

export function getStoredAuthToken() {
  if (typeof window === "undefined") {
    return "";
  }

  return (
    localStorage.getItem("authToken") ||
    localStorage.getItem("token") ||
    localStorage.getItem("accessToken") ||
    localStorage.getItem("jwt") ||
    ""
  );
}

export function clearStoredAuthToken() {
  localStorage.removeItem("authToken");
  localStorage.removeItem("token");
  localStorage.removeItem("accessToken");
  localStorage.removeItem("jwt");
}

export function getUserIdFromToken(token: string) {
  if (!token) {
    return "";
  }

  try {
    const [, payload] = token.split(".");
    if (!payload) {
      return "";
    }

    const normalized = payload.replace(/-/g, "+").replace(/_/g, "/");
    const padded = normalized.padEnd(
      normalized.length + ((4 - (normalized.length % 4)) % 4),
      "=",
    );
    const parsed = JSON.parse(atob(padded));
    return typeof parsed.userId === "string" ? parsed.userId : "";
  } catch {
    return "";
  }
}

export async function getCurrentUser() {
  const token = getStoredAuthToken();
  if (!token) {
    return null;
  }

  const response = await fetch(`${apiBase}/api/users/me`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  if (response.status === 401) {
    clearStoredAuthToken();
    return null;
  }
  if (!response.ok) {
    throw new Error("Could not load current user.");
  }

  return (await response.json()) as CurrentUser;
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
