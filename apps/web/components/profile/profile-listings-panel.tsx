import Link from "next/link";
import { Plus } from "lucide-react";
import { ProfileListingManagementCard } from "@/components/profile/profile-listing-management-card";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";

export function ProfileListingsPanel() {
  return (
    <section className={`${profilePanelClassName} flex flex-col`}>
      <div className="flex flex-col gap-5 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p className="text-xs font-bold uppercase tracking-[0.22em] text-brand-orange">
            Housing posts
          </p>
          <h1 className="mt-2 text-3xl font-bold tracking-tight text-gray-900">
            My Current Listings
          </h1>
        </div>

        <Link
          href="/create-post"
          className="inline-flex items-center justify-center gap-3 rounded-2xl border border-brand-orange bg-white px-7 py-4 text-xs font-bold uppercase tracking-[0.16em] text-brand-orange shadow-sm shadow-orange-900/5 transition-all hover:bg-brand-cream active:scale-95"
        >
          <Plus className="h-5 w-5" />
          Add Listing
        </Link>
      </div>

      <div className="mt-8 min-h-0 flex-1 space-y-5 overflow-y-auto pr-1">
        <ProfileListingManagementCard />
      </div>
    </section>
  );
}
