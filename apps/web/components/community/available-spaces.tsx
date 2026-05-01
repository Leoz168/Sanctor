import { SlidersHorizontal } from "lucide-react";
import { Button } from "@/components/ui/button";
import { PropertyCard } from "@/components/community/property-card";

const properties = [
  {
    href: "/communities/computer-science-collective/listings/harbord-room",
    image: "/images/listing-1.jpg",
    price: "$1,240/mo",
    title: "Harbord Village Sublet Room",
    location: "Downtown Toronto, 12 min to campus",
    beds: 1,
    baths: 1,
  },
  {
    href: "/communities/computer-science-collective/listings/annex-suite",
    image: "/images/listing-2.jpg",
    price: "$1,680/mo",
    title: "Annex Shared Suite",
    location: "Annex, near St. George Station",
    beds: 2,
    baths: 1,
  },
  {
    href: "/communities/computer-science-collective/listings/campus-loft",
    image: "/images/listing-3.jpg",
    price: "$2,050/mo",
    title: "Campus Loft Lease Transfer",
    location: "Spadina Corridor",
    beds: 2,
    baths: 2,
  },
];

export function AvailableSpaces() {
  return (
    <section className="px-4 py-12 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-6xl">
        <div className="mb-8 flex items-start justify-between">
          <div>
            <h2 className="mb-1 text-2xl font-bold text-foreground">
              Available Listings
            </h2>
            <p className="text-muted-foreground">
              Trusted housing posts currently active inside this community.
            </p>
          </div>
          <Button variant="outline" size="icon" className="rounded-lg">
            <SlidersHorizontal className="h-4 w-4" />
            <span className="sr-only">Filter listings</span>
          </Button>
        </div>

        <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3">
          {properties.map((property) => (
            <PropertyCard key={property.title} {...property} />
          ))}
        </div>
      </div>
    </section>
  );
}
