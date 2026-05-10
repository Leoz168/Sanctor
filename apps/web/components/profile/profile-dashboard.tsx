"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { Star } from "lucide-react";
import { ProfileAccountPanel } from "@/components/profile/profile-account-panel";
import { ProfileBookmarksPanel } from "@/components/profile/profile-bookmarks-panel";
import { ProfileCommunitiesPanel } from "@/components/profile/profile-communities-panel";
import {
  ProfileFormCard,
  type ProfileFormState,
} from "@/components/profile/profile-form-card";
import { ProfileListingsPanel } from "@/components/profile/profile-listings-panel";
import { ProfilePlaceholderPanel } from "@/components/profile/profile-placeholder-panel";
import {
  ProfileSidebar,
  type ProfileTab,
} from "@/components/profile/profile-sidebar";
import {
  clearStoredAuthToken,
  getStoredAuthToken,
} from "@/lib/auth-client";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

type Institution = {
  id: string;
  name: string;
};

type UserProfile = {
  id: string;
  username: string;
  email: string;
  avatar?: string;
  bio?: string;
  gender?: string;
  age?: number | null;
  institutionId?: string | null;
  major?: string | null;
};

const emptyFormState: ProfileFormState = {
  username: "",
  email: "",
  avatar: "",
  bio: "",
  gender: "",
  age: "",
  institutionId: "",
  major: "",
};

export function ProfileDashboard() {
  const [activeTab, setActiveTab] = useState<ProfileTab>("profile");
  const [form, setForm] = useState<ProfileFormState>(emptyFormState);
  const [institutions, setInstitutions] = useState<Institution[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);
  const [saveError, setSaveError] = useState<string | null>(null);
  const [saveStatus, setSaveStatus] = useState<string | null>(null);
  const [deleteError, setDeleteError] = useState<string | null>(null);
  const [isSaving, setIsSaving] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const router = useRouter();

  useEffect(() => {
    let isMounted = true;

    async function loadProfile() {
      const token = getStoredAuthToken();
      if (!token) {
        router.replace("/login");
        return;
      }

      setIsLoading(true);
      setLoadError(null);

      try {
        const headers = { Authorization: `Bearer ${token}` };
        const [profileResponse, institutionsResponse] = await Promise.all([
          fetch(`${apiBase}/api/users/me`, { headers }),
          fetch(`${apiBase}/api/institutions`, { headers }),
        ]);

        if (profileResponse.status === 401) {
          clearStoredAuthToken();
          router.replace("/login");
          return;
        }

        const profileData = (await profileResponse.json().catch(() => null)) as UserProfile | null;
        if (!profileResponse.ok || !profileData) {
          throw new Error("Unable to load your profile right now.");
        }

        const institutionData = institutionsResponse.ok
          ? ((await institutionsResponse.json().catch(() => [])) as Institution[])
          : [];

        if (!isMounted) {
          return;
        }

        setInstitutions(institutionData);
        setForm({
          username: profileData.username ?? "",
          email: profileData.email ?? "",
          avatar: profileData.avatar ?? "",
          bio: profileData.bio ?? "",
          gender: profileData.gender ?? "",
          age: profileData.age != null ? String(profileData.age) : "",
          institutionId: profileData.institutionId ?? "",
          major: profileData.major ?? "",
        });
      } catch (error) {
        if (isMounted) {
          setLoadError(error instanceof Error ? error.message : "Unable to load your profile.");
        }
      } finally {
        if (isMounted) {
          setIsLoading(false);
        }
      }
    }

    void loadProfile();

    return () => {
      isMounted = false;
    };
  }, [router]);

  const handleFieldChange = (field: keyof ProfileFormState, value: string) => {
    setForm((current) => ({
      ...current,
      [field]: value,
    }));
    setSaveError(null);
    setSaveStatus(null);
  };

  const handleSave = async () => {
    const token = getStoredAuthToken();
    if (!token) {
      clearStoredAuthToken();
      router.replace("/login");
      return;
    }

    setIsSaving(true);
    setSaveError(null);
    setSaveStatus(null);

    try {
      const response = await fetch(`${apiBase}/api/users/me`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          username: form.username.trim(),
          email: form.email.trim(),
          avatar: form.avatar.trim(),
          bio: form.bio.trim(),
          gender: form.gender || undefined,
          age: form.age.trim() ? Number(form.age) : null,
          institutionId: form.institutionId || "",
          major: form.major.trim() || null,
        }),
      });

      if (response.status === 401) {
        clearStoredAuthToken();
        router.replace("/login");
        return;
      }

      const data = (await response.json().catch(() => null)) as UserProfile | { error?: string } | null;
      if (!response.ok || !data || typeof (data as UserProfile).id !== "string") {
        const message =
          data && typeof data === "object" && "error" in data && typeof data.error === "string"
            ? data.error
            : "Unable to save your profile right now.";
        throw new Error(message);
      }

      const updatedProfile = data as UserProfile;
      setForm({
        username: updatedProfile.username ?? "",
        email: updatedProfile.email ?? "",
        avatar: updatedProfile.avatar ?? "",
        bio: updatedProfile.bio ?? "",
        gender: updatedProfile.gender ?? "",
        age: updatedProfile.age != null ? String(updatedProfile.age) : "",
        institutionId: updatedProfile.institutionId ?? "",
        major: updatedProfile.major ?? "",
      });
      setSaveStatus("Your profile has been updated.");
    } catch (error) {
      setSaveError(error instanceof Error ? error.message : "Unable to save your profile.");
    } finally {
      setIsSaving(false);
    }
  };

  const handleDelete = async () => {
    const token = getStoredAuthToken();
    if (!token) {
      clearStoredAuthToken();
      router.replace("/login");
      return;
    }

    if (!window.confirm("Delete your account permanently? This cannot be undone.")) {
      return;
    }

    setIsDeleting(true);
    setDeleteError(null);

    try {
      const response = await fetch(`${apiBase}/api/users/me`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok && response.status !== 204) {
        const message = await response.text();
        throw new Error(message || "Unable to delete your account.");
      }

      clearStoredAuthToken();
      router.replace("/login");
    } catch (error) {
      setDeleteError(error instanceof Error ? error.message : "Unable to delete your account.");
      setIsDeleting(false);
    }
  };

  if (isLoading) {
    return (
      <>
        <ProfileSidebar activeTab={activeTab} onTabChange={setActiveTab} />
        <ProfilePlaceholderPanel
          title="Loading profile"
          description="We’re pulling in your account details and institution settings now."
          icon={<Star className="h-6 w-6" />}
        />
      </>
    );
  }

  if (loadError) {
    return (
      <>
        <ProfileSidebar activeTab={activeTab} onTabChange={setActiveTab} />
        <ProfilePlaceholderPanel
          title="Profile unavailable"
          description={loadError}
          icon={<Star className="h-6 w-6" />}
        />
      </>
    );
  }

  return (
    <>
      <ProfileSidebar activeTab={activeTab} onTabChange={setActiveTab} />
      {activeTab === "profile" ? (
        <ProfileFormCard
          form={form}
          institutions={institutions}
          statusMessage={saveStatus ?? undefined}
          errorMessage={saveError ?? undefined}
          isSaving={isSaving}
          onFieldChange={handleFieldChange}
          onSave={handleSave}
        />
      ) : null}
      {activeTab === "listings" ? <ProfileListingsPanel /> : null}
      {activeTab === "communities" ? <ProfileCommunitiesPanel /> : null}
      {activeTab === "bookmarks" ? <ProfileBookmarksPanel /> : null}
      {activeTab === "account" ? (
        <ProfileAccountPanel
          email={form.email}
          username={form.username}
          isDeleting={isDeleting}
          errorMessage={deleteError ?? undefined}
          onDelete={handleDelete}
        />
      ) : null}
    </>
  );
}
