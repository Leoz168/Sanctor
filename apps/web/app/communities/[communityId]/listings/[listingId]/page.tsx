import Link from "next/link";
import { ArrowLeft, Bath, Bed, CalendarDays, Home } from "lucide-react";
import { AppShell } from "@/components/layout/app-shell";
import { ListingAmenitiesSection } from "@/components/listing/listing-amenities-section";
import { ListingCommentsSection } from "@/components/listing/listing-comments-section";
import { ListingInfoCard } from "@/components/listing/listing-info-card";
import { ListingMediaSection } from "@/components/listing/listing-media-section";

const listing = {
  communityName: "Computer Science Collective",
  title: "Modern Studio near campus",
  location: "St. George Campus, Toronto",
  verifiedLabel: "Community verified",
  monthlyRent: 1200,
  badge: "New",
  description:
    "A bright, furnished studio close to campus with a quiet study setup, warm natural light, and quick access to transit, grocery stops, and student spaces. The listing is posted through a community thread so students can ask questions, verify context, and keep housing conversations connected.",
  images: [
    { src: "/images/listing-1.jpg", alt: "Bright student studio bedroom" },
    { src: "/images/listing-2.jpg", alt: "Shared apartment kitchen" },
    { src: "/images/listing-3.jpg", alt: "Sunny apartment living area" },
    { src: "/images/listing-4.jpg", alt: "Quiet study and sleeping space" },
  ],
  stats: [
    { icon: Bed, label: "1 Bed" },
    { icon: Bath, label: "1 Bath" },
    { icon: Home, label: "Studio" },
    { icon: CalendarDays, label: "Fall term" },
  ],
  amenities: [
    "Furnished",
    "High-speed WiFi",
    "Laundry nearby",
    "Study desk",
    "Transit friendly",
    "Female friendly",
  ],
  poster: {
    name: "Maya K.",
    status: "Verified student poster",
    avatar: "/images/community-1.jpg",
    email: "maya.k@school.edu",
  },
};

export default function CommunityListingPage() {
  return (
    <AppShell surface="cream">
      <div className="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
        <Link
          href="/post-listings"
          className="mb-6 inline-flex items-center gap-2 text-sm font-bold text-gray-500 transition-colors hover:text-brand-orange"
        >
          <ArrowLeft size={18} />
          Back to listings
        </Link>

        <div className="grid items-stretch gap-8 lg:grid-cols-[1fr_380px]">
          <ListingMediaSection
            images={listing.images}
            description={listing.description}
          />
          <ListingInfoCard
            communityName={listing.communityName}
            title={listing.title}
            location={listing.location}
            verifiedLabel={listing.verifiedLabel}
            monthlyRent={listing.monthlyRent}
            badge={listing.badge}
            stats={listing.stats}
            poster={listing.poster}
          />
        </div>

        <div className="mt-8">
          <ListingAmenitiesSection amenities={listing.amenities} />
        </div>

        <div className="mt-6">
          <ListingCommentsSection />
        </div>
      </div>
    </AppShell>
  );
}
