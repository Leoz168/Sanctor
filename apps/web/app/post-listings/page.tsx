"use client";

import { useEffect, useState } from "react";
import { HousingFilterPanel } from "@/components/catalog/housing-filter-panel";
import {
  PaginatedListingsGrid,
  type PaginatedListing,
} from "@/components/catalog/paginated-listings-grid";
import { AppShell } from "@/components/layout/app-shell";
import { getStoredAuthToken, getUserIdFromToken } from "@/lib/auth-client";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
const fallbackImage = "/images/listing-1.jpg";

type BackendPost = {
  id: string;
  title?: string;
  address?: string;
  price?: number;
  rooms?: number;
  bathrooms?: number;
  created_at?: string;
};

type BackendPicture = {
  url?: string;
};

type BookmarkStatusResponse = {
  isBookmarked: boolean;
};

export default function PostListingsPage() {
  const [listings, setListings] = useState<PaginatedListing[]>([]);
  const [bookmarkedListingIds, setBookmarkedListingIds] = useState<Set<string>>(new Set());
  const [pendingBookmarkIds, setPendingBookmarkIds] = useState<Set<string>>(new Set());
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    let isMounted = true;

    async function loadListings() {
      setIsLoading(true);
      setError("");

      try {
        const response = await fetch(`${apiBase}/api/posts/search?limit=100`, {
          cache: "no-store",
        });
        if (!response.ok) {
          throw new Error("Could not load listings.");
        }

        const posts = dedupePostsByID((await response.json()) as BackendPost[]);
        const token = getStoredAuthToken();
        const userId = getUserIdFromToken(token);
        const mappedListings = await Promise.all(
          posts.map(async (post) => ({
            id: post.id,
            title: post.title || "Untitled listing",
            price: post.price ?? 0,
            location: post.address || "Address not provided",
            beds: post.rooms ?? 0,
            baths: post.bathrooms ?? 0,
            image: await loadPrimaryImage(post.id, token),
            badge: isRecent(post.created_at) ? ("new" as const) : undefined,
          })),
        );

        if (isMounted) {
          setListings(mappedListings);
          if (token && userId) {
            const bookmarkedIds = await loadBookmarkedListingIds(posts, token, userId);
            if (isMounted) {
              setBookmarkedListingIds(bookmarkedIds);
            }
          } else {
            setBookmarkedListingIds(new Set());
          }
        }
      } catch (loadError) {
        if (isMounted) {
          setError(loadError instanceof Error ? loadError.message : "Could not load listings.");
        }
      } finally {
        if (isMounted) {
          setIsLoading(false);
        }
      }
    }

    loadListings();

    return () => {
      isMounted = false;
    };
  }, []);

  return (
    <AppShell>
      <div className="mx-auto max-w-7xl px-4 pt-6 pb-10 sm:px-6 lg:px-8">
        <HousingFilterPanel />

        {isLoading && (
          <p className="py-12 text-center text-sm font-bold uppercase tracking-[0.18em] text-gray-400">
            Loading listings...
          </p>
        )}

        {!isLoading && error && (
          <p className="py-12 text-center text-sm font-bold text-red-600">{error}</p>
        )}

        {!isLoading && !error && listings.length === 0 && (
          <p className="py-12 text-center text-sm font-bold uppercase tracking-[0.18em] text-gray-400">
            No listings yet.
          </p>
        )}

        {!isLoading && !error && listings.length > 0 && (
          <PaginatedListingsGrid
            listings={listings}
            pageSize={20}
            bookmarkedListingIds={bookmarkedListingIds}
            pendingBookmarkIds={pendingBookmarkIds}
            onToggleBookmark={handleToggleBookmark}
          />
        )}
      </div>
    </AppShell>
  );

  async function handleToggleBookmark(listingId: string | number) {
    const postId = String(listingId);
    const token = getStoredAuthToken();
    const userId = getUserIdFromToken(token);
    if (!token || !userId) {
      if (window.confirm("Log in to save listings. Go to login now?")) {
        window.location.href = "/login";
      }
      return;
    }

    const isCurrentlyBookmarked = bookmarkedListingIds.has(postId);
    const confirmed = window.confirm(
      isCurrentlyBookmarked
        ? "Remove this listing from your saved listings?"
        : "Save this listing?",
    );
    if (!confirmed) {
      return;
    }

    setPendingBookmarkIds((current) => new Set(current).add(postId));
    try {
      if (isCurrentlyBookmarked) {
        const response = await fetch(
          `${apiBase}/api/posts/bookmarks/delete?userId=${userId}&postId=${postId}`,
          {
            method: "DELETE",
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        if (!response.ok) {
          throw new Error((await response.text()) || "Could not remove saved listing.");
        }

        setBookmarkedListingIds((current) => {
          const next = new Set(current);
          next.delete(postId);
          return next;
        });
      } else {
        const response = await fetch(`${apiBase}/api/posts/bookmarks/create`, {
          method: "POST",
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ userId, postId }),
        });
        if (!response.ok) {
          const message = await response.text();
          if (!message.toLowerCase().includes("already bookmarked")) {
            throw new Error(message || "Could not save listing.");
          }
        }

        setBookmarkedListingIds((current) => new Set(current).add(postId));
      }
    } catch (bookmarkError) {
      window.alert(bookmarkError instanceof Error ? bookmarkError.message : "Could not update saved listing.");
    } finally {
      setPendingBookmarkIds((current) => {
        const next = new Set(current);
        next.delete(postId);
        return next;
      });
    }
  }
}

async function loadPrimaryImage(postID: string, token: string | null) {
  if (!token) {
    return fallbackImage;
  }

  try {
    const response = await fetch(`${apiBase}/api/pictures?ownerType=post&ownerId=${postID}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (!response.ok) {
      return fallbackImage;
    }

    const pictures = (await response.json()) as BackendPicture[];
    return pictures[0]?.url || fallbackImage;
  } catch {
    return fallbackImage;
  }
}

function isRecent(value?: string) {
  if (!value) {
    return false;
  }

  const createdAt = new Date(value).getTime();
  if (Number.isNaN(createdAt)) {
    return false;
  }

  return Date.now() - createdAt < 1000 * 60 * 60 * 24 * 7;
}

function dedupePostsByID(posts: BackendPost[]) {
  const seen = new Set<string>();
  return posts.filter((post) => {
    if (!post.id || seen.has(post.id)) {
      return false;
    }

    seen.add(post.id);
    return true;
  });
}

async function loadBookmarkedListingIds(posts: BackendPost[], token: string, userId: string) {
  const statuses = await Promise.all(
    posts.map(async (post) => {
      try {
        const response = await fetch(
          `${apiBase}/api/posts/bookmarks/check?userId=${userId}&postId=${post.id}`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        );
        if (!response.ok) {
          return null;
        }

        const status = (await response.json()) as BookmarkStatusResponse;
        return status.isBookmarked ? post.id : null;
      } catch {
        return null;
      }
    }),
  );

  return new Set(statuses.filter((id): id is string => Boolean(id)));
}
