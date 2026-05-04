import Link from "next/link";
import { Plus, SlidersHorizontal } from "lucide-react";
import { CommunityListingPreviewCard } from "@/components/community/community-listing-preview-card";

interface CommunityListing {
  href: string;
  image: string;
  price: number;
  title: string;
  location: string;
  beds: number;
  baths: number;
  badge?: string;
}

interface CommunityListingsSectionProps {
  listings: CommunityListing[];
}

export function CommunityListingsSection({
  listings,
}: CommunityListingsSectionProps) {
  return (
    <section className="mt-8 rounded-[2rem] border border-gray-100 bg-white p-5 shadow-sm sm:p-6">
      <div className="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p className="mb-2 text-xs font-black uppercase tracking-[0.2em] text-brand-orange/70">
            Housing board
          </p>
          <h2 className="text-2xl font-bold text-gray-900">
            Available listings
          </h2>
        </div>

        <div className="flex gap-2">
          <button
            type="button"
            className="inline-flex h-11 w-11 items-center justify-center rounded-full border border-gray-100 bg-white text-gray-500 shadow-sm transition-colors hover:text-brand-orange"
            aria-label="Filter listings"
          >
            <SlidersHorizontal size={18} />
          </button>
          <Link
            href="/create-post"
            className="inline-flex items-center gap-2 rounded-full bg-brand-orange px-5 py-3 text-sm font-black text-white shadow-lg shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-[0.99]"
          >
            <Plus size={18} />
            Make a post
          </Link>
        </div>
      </div>

      <div className="grid gap-6 md:grid-cols-2 xl:grid-cols-3">
        {listings.map((listing) => (
          <CommunityListingPreviewCard key={listing.href} {...listing} />
        ))}
      </div>
    </section>
  );
}
