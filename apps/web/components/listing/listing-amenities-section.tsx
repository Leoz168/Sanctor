import { CheckCircle2 } from "lucide-react";

interface ListingAmenitiesSectionProps {
  amenities: string[];
}

export function ListingAmenitiesSection({
  amenities,
}: ListingAmenitiesSectionProps) {
  return (
    <section className="rounded-[2rem] border border-gray-100 bg-white p-5 shadow-sm sm:p-6">
      <h2 className="mb-4 text-xl font-bold text-gray-800">Amenities</h2>
      <div className="flex flex-wrap gap-2.5">
        {amenities.map((amenity) => (
          <div
            key={amenity}
            className="inline-flex items-center gap-2 rounded-full border border-orange-100 bg-orange-50/40 px-4 py-2.5"
          >
            <CheckCircle2 className="h-4 w-4 text-brand-orange" />
            <span className="text-sm font-bold text-gray-700">{amenity}</span>
          </div>
        ))}
      </div>
    </section>
  );
}
