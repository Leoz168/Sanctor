"use client";

import { useCallback, useRef, useState } from "react";
import { useRouter } from "next/navigation";

interface GoogleAuthButtonProps {
  label: string;
}

const apiBase = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";
const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || "";

type GoogleCredentialResponse = {
  credential?: string;
};

function getGoogle() {
  return (window as { google?: any }).google;
}

export function GoogleAuthButton({ label }: GoogleAuthButtonProps) {
  const router = useRouter();
  const initializedRef = useRef(false);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const ensureInitialized = useCallback(() => {
    const google = getGoogle();
    if (!google?.accounts?.id) {
      return false;
    }
    if (!initializedRef.current) {
      google.accounts.id.initialize({
        client_id: clientId,
        callback: (response: GoogleCredentialResponse) => {
          void handleCredential(response);
        },
      });
      initializedRef.current = true;
    }
    return true;
  }, []);

  const handleCredential = async (response: GoogleCredentialResponse) => {
    if (!response?.credential) {
      setError("Google sign-in failed. Try again.");
      return;
    }

    setLoading(true);
    setError("");
    try {
      const res = await fetch(`${apiBase}/api/auth/google`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ idToken: response.credential }),
      });

      const data = await res.json().catch(() => ({}));
      if (!res.ok) {
        setError(data?.error || "Google sign-in failed.");
        return;
      }

      if (data?.token) {
        localStorage.setItem("authToken", data.token);
      }
      router.push("/");
    } catch (err) {
      setError("Google sign-in failed. Try again.");
    } finally {
      setLoading(false);
    }
  };

  const handleClick = () => {
    if (!clientId) {
      setError("Google client ID is not configured.");
      return;
    }
    if (!ensureInitialized()) {
      setError("Google script not loaded yet.");
      return;
    }

    const google = getGoogle();
    google.accounts.id.prompt();
  };

  return (
    <div className="mt-3">
      <button
        type="button"
        onClick={handleClick}
        disabled={loading}
        className="flex w-full items-center justify-center gap-6 rounded-[0.2rem] border border-gray-300 bg-white px-5 py-3 text-base font-medium text-[#5f6368] shadow-md transition-all hover:bg-gray-50 active:scale-[0.99] disabled:cursor-not-allowed disabled:opacity-70 sm:text-lg"
      >
        {loading ? "Signing in..." : label}
      </button>
      {error && <p className="mt-2 text-sm text-red-600">{error}</p>}
    </div>
  );
}

function GoogleMark() {
  return (
    <svg aria-hidden="true" className="h-6 w-6 shrink-0" viewBox="0 0 48 48">
      <path
        fill="#EA4335"
        d="M24 9.5c3.54 0 6.71 1.22 9.21 3.6l6.85-6.85C35.9 2.38 30.47 0 24 0 14.62 0 6.51 5.38 2.56 13.22l7.98 6.19C12.43 13.72 17.74 9.5 24 9.5Z"
      />
      <path
        fill="#4285F4"
        d="M46.98 24.55c0-1.57-.15-3.09-.38-4.55H24v9.02h12.94c-.58 2.96-2.26 5.48-4.78 7.18l7.73 6c4.51-4.18 7.09-10.36 7.09-17.65Z"
      />
      <path
        fill="#FBBC05"
        d="M10.53 28.59A14.49 14.49 0 0 1 9.75 24c0-1.59.27-3.14.78-4.59l-7.98-6.19A23.94 23.94 0 0 0 0 24c0 3.88.93 7.56 2.56 10.78l7.97-6.19Z"
      />
      <path
        fill="#34A853"
        d="M24 48c6.48 0 11.93-2.13 15.89-5.81l-7.73-6c-2.15 1.45-4.92 2.3-8.16 2.3-6.26 0-11.57-4.22-13.47-9.91l-7.98 6.19C6.51 42.62 14.62 48 24 48Z"
      />
    </svg>
  );
}
