"use client";

import { FormEvent, useEffect, useState } from "react";
import { CommunityFilterPanel } from "@/components/catalog/community-filter-panel";
import { CommunityCard } from "@/components/community-card";
import { AppShell } from "@/components/layout/app-shell";

type ApiCommunity = {
  id: string;
  name: string;
  description?: string;
  isPrivate: boolean;
  createdBy?: string;
  createdAt?: string;
};

type CommunityCardData = {
  id: string;
  href: string;
  name: string;
  description: string;
  category: string;
  members: number;
  postsPerWeek: number;
  image?: string;
  isJoined: boolean;
  createdBy?: string;
  isOwner: boolean;
};

type Institution = {
  id: string;
  name: string;
  country?: string;
  region?: string;
};

type CreateCommunityForm = {
  name: string;
  description: string;
  institutionId: string;
  isPrivate: boolean;
};

type ManageCommunityForm = {
  id: string;
  name: string;
  description: string;
  isPrivate: boolean;
};

const fallbackImages = [
  "/images/community-1.jpg",
  "/images/community-4.jpg",
  "/images/community-5.jpg",
];

function slugify(value: string) {
  return value
    .toLowerCase()
    .trim()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "");
}

function getStoredToken() {
  if (typeof window === "undefined") {
    return "";
  }

  return (
    localStorage.getItem("token") ||
    localStorage.getItem("authToken") ||
    localStorage.getItem("accessToken") ||
    localStorage.getItem("jwt") ||
    ""
  );
}

function decodeBase64Url(value: string) {
  const normalized = value.replace(/-/g, "+").replace(/_/g, "/");
  const padded = normalized.padEnd(normalized.length + ((4 - (normalized.length % 4)) % 4), "=");
  return atob(padded);
}

function getUserIdFromToken(token: string) {
  if (!token) {
    return "";
  }

  try {
    const [, payload] = token.split(".");
    if (!payload) {
      return "";
    }

    const parsed = JSON.parse(decodeBase64Url(payload));
    return typeof parsed.userId === "string" ? parsed.userId : "";
  } catch {
    return "";
  }
}

export default function CommunitiesPage() {
  const [communities, setCommunities] = useState<CommunityCardData[]>([]);
  const [institutions, setInstitutions] = useState<Institution[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isInstitutionsLoading, setIsInstitutionsLoading] = useState(false);
  const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);
  const [isCreating, setIsCreating] = useState(false);
  const [isUpdating, setIsUpdating] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [error, setError] = useState("");
  const [createError, setCreateError] = useState("");
  const [manageError, setManageError] = useState("");
  const [modalMode, setModalMode] = useState<"create" | "manage">("create");
  const [form, setForm] = useState<CreateCommunityForm>({
    name: "",
    description: "",
    institutionId: "",
    isPrivate: false,
  });
  const [manageForm, setManageForm] = useState<ManageCommunityForm>({
    id: "",
    name: "",
    description: "",
    isPrivate: false,
  });

  const refreshCommunities = async () => {
    const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
    const token = getStoredToken();
    const currentUserId = getUserIdFromToken(token);

    if (!token) {
      setError("Log in first so we can load your communities.");
      setIsLoading(false);
      return;
    }

    const response = await fetch(`${apiBase}/api/communities`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      const message = await response.text();
      throw new Error(message || "Failed to load communities.");
    }

    const data: ApiCommunity[] = await response.json();
    const mapped = data.map((community, index) => ({
      id: community.id,
      href: `/communities/${slugify(community.name) || community.id}`,
      name: community.name,
      description: community.description?.trim() || "No description yet.",
      category: community.isPrivate ? "PRIVATE" : "PUBLIC",
      members: 0,
      postsPerWeek: 0,
      image: fallbackImages[index % fallbackImages.length],
      isJoined: false,
      createdBy: community.createdBy,
      isOwner: Boolean(currentUserId) && community.createdBy === currentUserId,
    }));

    setCommunities(mapped);

    setManageForm((current) => {
      const ownedCommunities = mapped.filter((community) => community.isOwner);
      if (ownedCommunities.length === 0) {
        return { id: "", name: "", description: "", isPrivate: false };
      }

      const selected = ownedCommunities.find((community) => community.id === current.id) || ownedCommunities[0];
      return {
        id: selected.id,
        name: selected.name,
        description: selected.description === "No description yet." ? "" : selected.description,
        isPrivate: selected.category === "PRIVATE",
      };
    });
  };

  useEffect(() => {
    const load = async () => {
      try {
        await refreshCommunities();
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load communities.");
      } finally {
        setIsLoading(false);
      }
    };

    load();
  }, []);

  const loadInstitutions = async () => {
    if (institutions.length > 0 || isInstitutionsLoading) {
      return;
    }

    const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
    const token = getStoredToken();

    if (!token) {
      setCreateError("Log in first so we can create a community.");
      return;
    }

    setIsInstitutionsLoading(true);
    setCreateError("");

    try {
      const response = await fetch(`${apiBase}/api/institutions`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Failed to load institutions.");
      }

      const data: Institution[] = await response.json();
      setInstitutions(data);
      setForm((current) => ({
        ...current,
        institutionId: current.institutionId || data[0]?.id || "",
      }));
    } catch (err) {
      setCreateError(err instanceof Error ? err.message : "Failed to load institutions.");
    } finally {
      setIsInstitutionsLoading(false);
    }
  };

  const openCreateModal = async () => {
    setCreateError("");
    setManageError("");
    setIsCreateModalOpen(true);
    setModalMode("create");
    await loadInstitutions();
  };

  const closeCreateModal = () => {
    if (isCreating || isUpdating || isDeleting) {
      return;
    }

    setIsCreateModalOpen(false);
    setCreateError("");
    setManageError("");
  };

  const handleCreateCommunity = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const token = getStoredToken();
    const createdBy = getUserIdFromToken(token);
    const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

    if (!token || !createdBy) {
      setCreateError("Log in again so we can verify your account before creating a community.");
      return;
    }

    if (!form.name.trim()) {
      setCreateError("Give your community a name first.");
      return;
    }

    if (!form.institutionId) {
      setCreateError("Choose an institution before creating the community.");
      return;
    }

    setIsCreating(true);
    setCreateError("");

    try {
      const response = await fetch(`${apiBase}/api/communities/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          name: form.name.trim(),
          description: form.description.trim(),
          isPrivate: form.isPrivate,
          institutionId: form.institutionId,
          createdBy,
        }),
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Failed to create community.");
      }

      await refreshCommunities();
      setForm({
        name: "",
        description: "",
        institutionId: institutions[0]?.id || "",
        isPrivate: false,
      });
      setIsCreateModalOpen(false);
    } catch (err) {
      setCreateError(err instanceof Error ? err.message : "Failed to create community.");
    } finally {
      setIsCreating(false);
    }
  };

  const ownedCommunities = communities.filter((community) => community.isOwner);

  const syncManagedCommunity = (communityID: string) => {
    const selected = ownedCommunities.find((community) => community.id === communityID);
    if (!selected) {
      setManageForm({ id: "", name: "", description: "", isPrivate: false });
      return;
    }

    setManageForm({
      id: selected.id,
      name: selected.name,
      description: selected.description === "No description yet." ? "" : selected.description,
      isPrivate: selected.category === "PRIVATE",
    });
  };

  const handleUpdateCommunity = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const token = getStoredToken();
    const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

    if (!token) {
      setManageError("Log in again so we can verify your account before updating a community.");
      return;
    }

    if (!manageForm.id) {
      setManageError("Choose one of your communities first.");
      return;
    }

    if (!manageForm.name.trim()) {
      setManageError("Give your community a name first.");
      return;
    }

    setIsUpdating(true);
    setManageError("");

    try {
      const response = await fetch(`${apiBase}/api/communities/update?id=${encodeURIComponent(manageForm.id)}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          name: manageForm.name.trim(),
          description: manageForm.description.trim(),
          isPrivate: manageForm.isPrivate,
        }),
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Failed to update community.");
      }

      await refreshCommunities();
    } catch (err) {
      setManageError(err instanceof Error ? err.message : "Failed to update community.");
    } finally {
      setIsUpdating(false);
    }
  };

  const handleDeleteCommunity = async () => {
    const token = getStoredToken();
    const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

    if (!token) {
      setManageError("Log in again so we can verify your account before deleting a community.");
      return;
    }

    if (!manageForm.id) {
      setManageError("Choose one of your communities first.");
      return;
    }

    const confirmed = window.confirm(`Delete "${manageForm.name}"? This cannot be undone.`);
    if (!confirmed) {
      return;
    }

    setIsDeleting(true);
    setManageError("");

    try {
      const response = await fetch(`${apiBase}/api/communities/delete?id=${encodeURIComponent(manageForm.id)}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        const message = await response.text();
        throw new Error(message || "Failed to delete community.");
      }

      await refreshCommunities();

      const remainingOwnedCommunities = ownedCommunities.filter((community) => community.id !== manageForm.id);
      if (remainingOwnedCommunities.length === 0) {
        setModalMode("create");
        setManageForm({ id: "", name: "", description: "", isPrivate: false });
      } else {
        syncManagedCommunity(remainingOwnedCommunities[0].id);
      }
    } catch (err) {
      setManageError(err instanceof Error ? err.message : "Failed to delete community.");
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <AppShell>
      <div className="mx-auto max-w-7xl px-4 pb-10 sm:px-6 lg:px-8">
        <CommunityFilterPanel onCreateCommunity={openCreateModal} />

        {isLoading ? (
          <div className="rounded-[2rem] border border-gray-100 bg-white p-10 text-center text-gray-500 shadow-sm">
            Loading communities...
          </div>
        ) : error ? (
          <div className="rounded-[2rem] border border-red-100 bg-red-50 p-10 text-center text-red-700 shadow-sm">
            {error}
          </div>
        ) : communities.length === 0 ? (
          <div className="rounded-[2rem] border border-gray-100 bg-white p-10 text-center text-gray-500 shadow-sm">
            No communities are available yet.
          </div>
        ) : (
          <div className="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
            {communities.map((community) => (
              <CommunityCard key={community.id} {...community} />
            ))}
          </div>
        )}
      </div>

      {isCreateModalOpen ? (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-gray-900/40 px-4 py-8 backdrop-blur-sm">
          <div className="w-full max-w-2xl rounded-[2rem] bg-white p-8 shadow-2xl">
            <div className="flex items-start justify-between gap-4">
              <div>
                <p className="text-sm font-bold uppercase tracking-[0.2em] text-brand-orange">
                  {modalMode === "create" ? "Create" : "Manage"}
                </p>
                <h2 className="mt-2 text-3xl font-bold text-gray-900">
                  {modalMode === "create" ? "Start a new community" : "Update or delete your community"}
                </h2>
                <p className="mt-2 text-sm text-gray-500">
                  {modalMode === "create"
                    ? "This will create the community in Sanctor and make you its owner."
                    : "These actions call the protected community update and delete backend endpoints."}
                </p>
              </div>
              <button
                type="button"
                onClick={closeCreateModal}
                className="rounded-full border border-gray-200 px-4 py-2 text-sm font-semibold text-gray-500 transition-colors hover:border-gray-300 hover:text-gray-900"
              >
                Close
              </button>
            </div>

            <div className="mt-6 inline-flex rounded-full border border-gray-200 bg-gray-50 p-1">
              <button
                type="button"
                onClick={() => setModalMode("create")}
                className={`rounded-full px-4 py-2 text-sm font-semibold transition-colors ${
                  modalMode === "create" ? "bg-white text-gray-900 shadow-sm" : "text-gray-500 hover:text-gray-900"
                }`}
              >
                Create
              </button>
              <button
                type="button"
                onClick={() => {
                  setModalMode("manage");
                  if (ownedCommunities.length > 0 && !manageForm.id) {
                    syncManagedCommunity(ownedCommunities[0].id);
                  }
                }}
                className={`rounded-full px-4 py-2 text-sm font-semibold transition-colors ${
                  modalMode === "manage" ? "bg-white text-gray-900 shadow-sm" : "text-gray-500 hover:text-gray-900"
                }`}
              >
                Manage
              </button>
            </div>

            {modalMode === "create" ? (
            <form className="mt-8 space-y-6" onSubmit={handleCreateCommunity}>
              <div className="grid gap-6 md:grid-cols-2">
                <label className="block">
                  <span className="mb-2 block text-sm font-semibold text-gray-700">Community name</span>
                  <input
                    type="text"
                    value={form.name}
                    onChange={(event) =>
                      setForm((current) => ({
                        ...current,
                        name: event.target.value,
                      }))
                    }
                    placeholder="U of T Off-Campus Housing"
                    className="w-full rounded-2xl border border-gray-200 bg-white px-4 py-3 text-sm text-gray-900 outline-none transition focus:border-brand-orange focus:ring-2 focus:ring-brand-orange/20"
                  />
                </label>

                <label className="block">
                  <span className="mb-2 block text-sm font-semibold text-gray-700">Institution</span>
                  <select
                    value={form.institutionId}
                    onChange={(event) =>
                      setForm((current) => ({
                        ...current,
                        institutionId: event.target.value,
                      }))
                    }
                    disabled={isInstitutionsLoading}
                    className="w-full rounded-2xl border border-gray-200 bg-white px-4 py-3 text-sm text-gray-900 outline-none transition focus:border-brand-orange focus:ring-2 focus:ring-brand-orange/20 disabled:bg-gray-50"
                  >
                    <option value="">
                      {isInstitutionsLoading ? "Loading institutions..." : "Select an institution"}
                    </option>
                    {institutions.map((institution) => (
                      <option key={institution.id} value={institution.id}>
                        {institution.name}
                      </option>
                    ))}
                  </select>
                </label>
              </div>

              <label className="block">
                <span className="mb-2 block text-sm font-semibold text-gray-700">Description</span>
                <textarea
                  value={form.description}
                  onChange={(event) =>
                    setForm((current) => ({
                      ...current,
                      description: event.target.value,
                    }))
                  }
                  rows={5}
                  placeholder="What is this community for, and who should join?"
                  className="w-full rounded-3xl border border-gray-200 bg-white px-4 py-4 text-sm text-gray-900 outline-none transition focus:border-brand-orange focus:ring-2 focus:ring-brand-orange/20"
                />
              </label>

              <label className="flex items-center gap-3 rounded-2xl border border-gray-200 bg-brand-cream/60 px-4 py-4">
                <input
                  type="checkbox"
                  checked={form.isPrivate}
                  onChange={(event) =>
                    setForm((current) => ({
                      ...current,
                      isPrivate: event.target.checked,
                    }))
                  }
                  className="h-4 w-4 rounded border-gray-300 text-brand-orange focus:ring-brand-orange/30"
                />
                <span>
                  <span className="block text-sm font-semibold text-gray-900">Private community</span>
                  <span className="block text-sm text-gray-500">
                    New members will need approval before they can join.
                  </span>
                </span>
              </label>

              {createError ? (
                <div className="rounded-2xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-700">
                  {createError}
                </div>
              ) : null}

              <div className="flex flex-col-reverse gap-3 sm:flex-row sm:justify-end">
                <button
                  type="button"
                  onClick={closeCreateModal}
                  className="rounded-full border border-gray-200 px-5 py-3 text-sm font-semibold text-gray-600 transition-colors hover:border-gray-300 hover:text-gray-900"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  disabled={isCreating || isInstitutionsLoading}
                  className="rounded-full bg-brand-orange px-5 py-3 text-sm font-semibold text-white shadow-lg shadow-brand-orange/20 transition-all hover:bg-orange-600 disabled:cursor-not-allowed disabled:opacity-60"
                >
                  {isCreating ? "Creating..." : "Create community"}
                </button>
              </div>
            </form>
            ) : (
              <form className="mt-8 space-y-6" onSubmit={handleUpdateCommunity}>
                {ownedCommunities.length === 0 ? (
                  <div className="rounded-2xl border border-gray-100 bg-gray-50 px-4 py-4 text-sm text-gray-600">
                    You do not own any communities yet, so there is nothing to update or delete here.
                  </div>
                ) : (
                  <>
                    <label className="block">
                      <span className="mb-2 block text-sm font-semibold text-gray-700">Your communities</span>
                      <select
                        value={manageForm.id}
                        onChange={(event) => syncManagedCommunity(event.target.value)}
                        className="w-full rounded-2xl border border-gray-200 bg-white px-4 py-3 text-sm text-gray-900 outline-none transition focus:border-brand-orange focus:ring-2 focus:ring-brand-orange/20"
                      >
                        {ownedCommunities.map((community) => (
                          <option key={community.id} value={community.id}>
                            {community.name}
                          </option>
                        ))}
                      </select>
                    </label>

                    <label className="block">
                      <span className="mb-2 block text-sm font-semibold text-gray-700">Community name</span>
                      <input
                        type="text"
                        value={manageForm.name}
                        onChange={(event) =>
                          setManageForm((current) => ({
                            ...current,
                            name: event.target.value,
                          }))
                        }
                        className="w-full rounded-2xl border border-gray-200 bg-white px-4 py-3 text-sm text-gray-900 outline-none transition focus:border-brand-orange focus:ring-2 focus:ring-brand-orange/20"
                      />
                    </label>

                    <label className="block">
                      <span className="mb-2 block text-sm font-semibold text-gray-700">Description</span>
                      <textarea
                        value={manageForm.description}
                        onChange={(event) =>
                          setManageForm((current) => ({
                            ...current,
                            description: event.target.value,
                          }))
                        }
                        rows={5}
                        className="w-full rounded-3xl border border-gray-200 bg-white px-4 py-4 text-sm text-gray-900 outline-none transition focus:border-brand-orange focus:ring-2 focus:ring-brand-orange/20"
                      />
                    </label>

                    <label className="flex items-center gap-3 rounded-2xl border border-gray-200 bg-brand-cream/60 px-4 py-4">
                      <input
                        type="checkbox"
                        checked={manageForm.isPrivate}
                        onChange={(event) =>
                          setManageForm((current) => ({
                            ...current,
                            isPrivate: event.target.checked,
                          }))
                        }
                        className="h-4 w-4 rounded border-gray-300 text-brand-orange focus:ring-brand-orange/30"
                      />
                      <span>
                        <span className="block text-sm font-semibold text-gray-900">Private community</span>
                        <span className="block text-sm text-gray-500">
                          This maps directly to the update endpoint&apos;s `isPrivate` field.
                        </span>
                      </span>
                    </label>
                  </>
                )}

                {manageError ? (
                  <div className="rounded-2xl border border-red-100 bg-red-50 px-4 py-3 text-sm text-red-700">
                    {manageError}
                  </div>
                ) : null}

                <div className="flex flex-col-reverse gap-3 sm:flex-row sm:justify-between">
                  <button
                    type="button"
                    onClick={handleDeleteCommunity}
                    disabled={ownedCommunities.length === 0 || isUpdating || isDeleting}
                    className="rounded-full border border-red-200 px-5 py-3 text-sm font-semibold text-red-600 transition-colors hover:bg-red-50 disabled:cursor-not-allowed disabled:opacity-60"
                  >
                    {isDeleting ? "Deleting..." : "Delete community"}
                  </button>
                  <button
                    type="submit"
                    disabled={ownedCommunities.length === 0 || isUpdating || isDeleting}
                    className="rounded-full bg-brand-orange px-5 py-3 text-sm font-semibold text-white shadow-lg shadow-brand-orange/20 transition-all hover:bg-orange-600 disabled:cursor-not-allowed disabled:opacity-60"
                  >
                    {isUpdating ? "Saving..." : "Save changes"}
                  </button>
                </div>
              </form>
            )}
          </div>
        </div>
      ) : null}
    </AppShell>
  );
}
