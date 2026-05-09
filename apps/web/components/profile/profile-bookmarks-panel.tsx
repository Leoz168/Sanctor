"use client";

import { useEffect, useState } from "react";
import { Bookmark } from "lucide-react";
import { ListingCard } from "@/components/listing-card";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";
import { getStoredAuthToken, getUserIdFromToken } from "@/lib/auth-client";

const apiBase = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";
const fallbackImage = "/images/listing-1.jpg";

type BookmarkedPost = {
  id: string;
  title?: string;
  address?: string;
  price?: number;
  rooms?: number;
  bathrooms?: number;
};

type BackendPicture = {
  url?: string;
};

type SavedListing = {
  id: string;
  title: string;
  price: number;
  location: string;
  beds: number;
  baths: number;
  image: string;
};

export function ProfileBookmarksPanel() {
  const [listings, setListings] = useState<SavedListing[]>([]);
  const [pendingListingIds, setPendingListingIds] = useState<Set<string>>(new Set());
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    let isMounted = true;

    async function loadBookmarks() {
      setIsLoading(true);
      setError("");

      const token = getStoredAuthToken();
      const userId = getUserIdFromToken(token);
      if (!token || !userId) {
        setListings([]);
        setError("Log in to view your saved listings.");
        setIsLoading(false);
        return;
      }

      try {
        const response = await fetch(`${apiBase}/api/posts/bookmarks?userId=${userId}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        if (!response.ok) {
          throw new Error((await response.text()) || "Could not load saved listings.");
        }

        const posts = dedupePostsByID((await response.json()) as BookmarkedPost[]);
        const savedListings = await Promise.all(
          posts.map(async (post) => ({
            id: post.id,
            title: post.title || "Untitled listing",
            price: post.price ?? 0,
            location: post.address || "Address not provided",
            beds: post.rooms ?? 0,
            baths: post.bathrooms ?? 0,
            image: await loadPrimaryImage(post.id, token),
          })),
        );

        if (isMounted) {
          setListings(savedListings);
        }
      } catch (loadError) {
        if (isMounted) {
          setError(loadError instanceof Error ? loadError.message : "Could not load saved listings.");
        }
      } finally {
        if (isMounted) {
          setIsLoading(false);
        }
      }
    }

    loadBookmarks();

    return () => {
      isMounted = false;
    };
  }, []);

  async function handleRemoveBookmark(listingId: string | number) {
    const postId = String(listingId);
    const token = getStoredAuthToken();
    const userId = getUserIdFromToken(token);
    if (!token || !userId) {
      setError("Log in to manage your saved listings.");
      return;
    }
    if (!window.confirm("Remove this listing from your saved listings?")) {
      return;
    }

    setPendingListingIds((current) => new Set(current).add(postId));
    try {
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

      setListings((current) => current.filter((listing) => listing.id !== postId));
    } catch (removeError) {
      window.alert(removeError instanceof Error ? removeError.message : "Could not remove saved listing.");
    } finally {
      setPendingListingIds((current) => {
        const next = new Set(current);
        next.delete(postId);
        return next;
      });
    }
  }

  return (
    <section className={`${profilePanelClassName} flex flex-col`}>
      <div className="flex items-start justify-between gap-6">
        <div>
          <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-brand-cream text-brand-orange">
            <Bookmark className="h-6 w-6" />
          </div>
          <h1 className="mt-6 text-3xl font-bold tracking-tight text-gray-900">
            Bookmarked Listings
          </h1>
          <p className="mt-3 max-w-2xl text-base font-semibold leading-7 text-gray-500">
            Your saved listings are collected here so you can compare them later.
          </p>
        </div>
      </div>

      {isLoading ? (
        <p className="py-12 text-sm font-bold uppercase tracking-[0.18em] text-gray-400">
          Loading saved listings...
        </p>
      ) : null}

      {!isLoading && error ? (
        <p className="py-12 text-sm font-bold text-red-600">{error}</p>
      ) : null}

      {!isLoading && !error && listings.length === 0 ? (
        <p className="py-12 text-sm font-bold uppercase tracking-[0.18em] text-gray-400">
          No saved listings yet.
        </p>
      ) : null}

      {!isLoading && !error && listings.length > 0 ? (
        <div className="mt-8 grid grid-cols-1 gap-8 xl:grid-cols-2">
          {listings.map((listing) => (
            <ListingCard
              key={listing.id}
              {...listing}
              isBookmarked
              isBookmarkPending={pendingListingIds.has(listing.id)}
              onToggleBookmark={handleRemoveBookmark}
            />
          ))}
        </div>
      ) : null}
    </section>
  );
}

async function loadPrimaryImage(postId: string, token: string) {
  try {
    const response = await fetch(`${apiBase}/api/pictures?ownerType=post&ownerId=${postId}`, {
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

function dedupePostsByID(posts: BookmarkedPost[]) {
  const seen = new Set<string>();
  return posts.filter((post) => {
    if (!post.id || seen.has(post.id)) {
      return false;
    }

    seen.add(post.id);
    return true;
  });
}
