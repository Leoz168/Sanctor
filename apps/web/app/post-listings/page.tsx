import { Plus } from "lucide-react";
import { FloatingActionButton } from "@/components/catalog/floating-action-button";
import { HousingFilterPanel } from "@/components/catalog/housing-filter-panel";
import { AppShell } from "@/components/layout/app-shell";
import { ListingCard } from "@/components/listing-card";

const listings = [
  {
    id: 1,
    title: "Modern Studio near campus",
    price: 1200,
    location: "St. George Campus, Toronto",
    beds: 1,
    baths: 1,
    image: "/images/listing-1.jpg",
    badge: "featured" as const,
  },
  {
    id: 2,
    title: "Shared 3BR House - Female only",
    price: 850,
    location: "North Campus Area",
    beds: 3,
    baths: 2,
    image: "/images/listing-2.jpg",
    badge: "new" as const,
  },
  {
    id: 3,
    title: "Luxury Apartment in Downtown",
    price: 2100,
    location: "Downtown Core",
    beds: 2,
    baths: 2,
    image: "/images/listing-3.jpg",
  },
  {
    id: 4,
    title: "Cozy Loft for Students",
    price: 950,
    location: "East Side Campus",
    beds: 1,
    baths: 1,
    image: "/images/listing-4.jpg",
  },
  {
    id: 5,
    title: "Renovated Basement Suite",
    price: 1100,
    location: "West Campus Gardens",
    beds: 1,
    baths: 1,
    image: "/images/listing-5.jpg",
  },
  {
    id: 6,
    title: "Large 4BR Student Residence",
    price: 700,
    location: "Campus South",
    beds: 4,
    baths: 3,
    image: "/images/listing-6.jpg",
    badge: "new" as const,
  },
];

export default function PostListingsPage() {
  return (
    <AppShell floatingAction={<FloatingActionButton icon={Plus} href="/create-post">Make a post</FloatingActionButton>}>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
        <HousingFilterPanel />

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {listings.map((listing) => (
            <ListingCard key={listing.id} {...listing} />
          ))}
        </div>
      </div>
    </AppShell>
  );
}
