import Link from "next/link";
import { Plus } from "lucide-react";
import { ProfileListingManagementCard } from "@/components/profile/profile-listing-management-card";
import { profilePanelClassName } from "@/components/profile/profile-panel-styles";

export function ProfileListingsPanel() {
  return (
    <section className={`${profilePanelClassName} flex flex-col`}>
      <div>
        <p className="text-xs font-bold uppercase tracking-[0.22em] text-brand-orange">
          Housing posts
        </p>
        <h1 className="mt-2 text-3xl font-bold tracking-tight text-gray-900">
          My Current Listings
        </h1>
      </div>

      <div className="mt-8 min-h-0 flex-1 space-y-5 overflow-y-auto pr-1">
        <ProfileListingManagementCard />
      </div>

      <div className="mt-6 flex justify-end">
        <Link
          href="/create-post"
          className="inline-flex h-14 items-center justify-center gap-3 rounded-2xl bg-brand-orange px-7 text-sm font-bold uppercase tracking-[0.14em] text-white shadow-xl shadow-brand-orange/25 transition-all hover:bg-orange-600 active:scale-95"
        >
          <Plus className="h-5 w-5" />
          Add Listing
        </Link>
      </div>
    </section>
  );
}
