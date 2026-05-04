"use client";

import { useMemo, useState } from "react";
import type { ReactNode } from "react";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { ListingCard } from "@/components/listing-card";

export interface PaginatedListing {
  id: number;
  title: string;
  price: number;
  location: string;
  beds: number;
  baths: number;
  image: string;
  badge?: "featured" | "new";
}

interface PaginatedListingsGridProps {
  listings: PaginatedListing[];
  pageSize?: number;
}

export function PaginatedListingsGrid({
  listings,
  pageSize = 20,
}: PaginatedListingsGridProps) {
  const [currentPage, setCurrentPage] = useState(1);
  const totalPages = Math.max(1, Math.ceil(listings.length / pageSize));
  const pageStart = (currentPage - 1) * pageSize;
  const currentListings = listings.slice(pageStart, pageStart + pageSize);
  const visiblePages = useMemo(
    () => getVisiblePages(currentPage, totalPages),
    [currentPage, totalPages],
  );

  const goToPage = (page: number) => {
    const nextPage = Math.min(Math.max(1, page), totalPages);

    setCurrentPage(nextPage);
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

  return (
    <section aria-label="Housing listings">
      <div className="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
        {currentListings.map((listing) => (
          <ListingCard key={listing.id} {...listing} />
        ))}
      </div>

      {totalPages > 1 && (
        <nav
          className="flex items-center justify-center gap-3 py-12"
          aria-label="Listings pagination"
        >
          <PaginationButton
            onClick={() => goToPage(currentPage - 1)}
            disabled={currentPage === 1}
            ariaLabel="Previous listings page"
          >
            <ChevronLeft className="h-4 w-4" />
          </PaginationButton>

          {visiblePages.map((page, index) =>
            typeof page === "number" ? (
              <PaginationButton
                key={page}
                onClick={() => goToPage(page)}
                isActive={page === currentPage}
                ariaLabel={`Go to listings page ${page}`}
              >
                {page}
              </PaginationButton>
            ) : (
              <span
                key={`ellipsis-${index}`}
                className="flex h-9 w-9 items-center justify-center text-sm font-semibold text-gray-500"
              >
                ...
              </span>
            ),
          )}

          <PaginationButton
            onClick={() => goToPage(currentPage + 1)}
            disabled={currentPage === totalPages}
            ariaLabel="Next listings page"
          >
            <ChevronRight className="h-4 w-4" />
          </PaginationButton>
        </nav>
      )}
    </section>
  );
}

interface PaginationButtonProps {
  children: ReactNode;
  onClick: () => void;
  ariaLabel: string;
  disabled?: boolean;
  isActive?: boolean;
}

function PaginationButton({
  children,
  onClick,
  ariaLabel,
  disabled = false,
  isActive = false,
}: PaginationButtonProps) {
  return (
    <button
      type="button"
      onClick={onClick}
      disabled={disabled}
      aria-label={ariaLabel}
      aria-current={isActive ? "page" : undefined}
      className={`flex h-9 min-w-9 items-center justify-center rounded-full px-3 text-sm font-bold transition-all ${
        isActive
          ? "bg-gray-900 text-white shadow-md"
          : "text-gray-700 hover:bg-brand-cream hover:text-brand-orange"
      } disabled:pointer-events-none disabled:text-gray-300`}
    >
      {children}
    </button>
  );
}

function getVisiblePages(currentPage: number, totalPages: number) {
  if (totalPages <= 7) {
    return Array.from({ length: totalPages }, (_, index) => index + 1);
  }

  if (currentPage <= 4) {
    return [1, 2, 3, 4, "ellipsis", totalPages];
  }

  if (currentPage >= totalPages - 3) {
    return [
      1,
      "ellipsis",
      totalPages - 3,
      totalPages - 2,
      totalPages - 1,
      totalPages,
    ];
  }

  return [
    1,
    "ellipsis",
    currentPage - 1,
    currentPage,
    currentPage + 1,
    "ellipsis",
    totalPages,
  ];
}
