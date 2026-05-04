import Link from "next/link";
import { ArrowLeft } from "lucide-react";
import { AppShell } from "@/components/layout/app-shell";
import { CommunityAboutPanel } from "@/components/community/community-about-panel";
import { CommunityDetailHero } from "@/components/community/community-detail-hero";
import { CommunityListingsSection } from "@/components/community/community-listings-section";

const community = {
  name: "Computer Science Collective",
  category: "Academic community",
  image: "/images/community-1.jpg",
  members: 1248,
  postsPerWeek: 382,
  description:
    "A high-signal student space for course support, hackathon teams, housing leads, and everyday campus problem-solving. Students use this community to share trusted listings, find roommates, compare classes, and keep campus knowledge in one place.",
  listings: [
    {
      href: "/communities/computer-science-collective/listings/modern-studio-near-campus",
      image: "/images/listing-1.jpg",
      price: 1200,
      title: "Modern Studio near campus",
      location: "St. George Campus, Toronto",
      beds: 1,
      baths: 1,
      badge: "Featured",
    },
    {
      href: "/communities/computer-science-collective/listings/annex-shared-suite",
      image: "/images/listing-2.jpg",
      price: 850,
      title: "Annex Shared Suite",
      location: "Annex, near St. George Station",
      beds: 3,
      baths: 2,
      badge: "New",
    },
    {
      href: "/communities/computer-science-collective/listings/campus-loft-transfer",
      image: "/images/listing-3.jpg",
      price: 2100,
      title: "Campus Loft Lease Transfer",
      location: "Spadina Corridor",
      beds: 2,
      baths: 2,
    },
  ],
};

export default function CommunityPage() {
  return (
    <AppShell surface="cream">
      <div className="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
        <Link
          href="/communities"
          className="mb-6 inline-flex items-center gap-2 text-sm font-bold text-gray-500 transition-colors hover:text-brand-orange"
        >
          <ArrowLeft size={18} />
          Back to communities
        </Link>

        <CommunityDetailHero
          name={community.name}
          category={community.category}
          description={community.description}
          image={community.image}
          members={community.members}
          postsPerWeek={community.postsPerWeek}
        />

        <CommunityAboutPanel description={community.description} />
        <CommunityListingsSection listings={community.listings} />
      </div>
    </AppShell>
  );
}
