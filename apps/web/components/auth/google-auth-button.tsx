"use client";

import { useRef, useState } from "react";
import { useRouter } from "next/navigation";

interface GoogleAuthButtonProps {
  label: string;
}

const apiBase = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";
const clientId = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || "";

type GoogleCredentialResponse = {
  credential?: string;
};

type GoogleIdentityClient = {
  initialize: (config: {
    client_id: string;
    callback: (response: GoogleCredentialResponse) => void;
  }) => void;
  prompt: () => void;
};

type GoogleWindow = {
  google?: {
    accounts?: {
      id?: GoogleIdentityClient;
    };
  };
};

function getGoogle() {
  return (window as GoogleWindow).google;
}

export function GoogleAuthButton({ label }: GoogleAuthButtonProps) {
  const router = useRouter();
  const initializedRef = useRef(false);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

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
    } catch {
      setError("Google sign-in failed. Try again.");
    } finally {
      setLoading(false);
    }
  };

  const ensureInitialized = () => {
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
    google?.accounts?.id?.prompt();
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
