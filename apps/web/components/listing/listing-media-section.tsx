import { ListingPhotoCarousel } from "@/components/listing/listing-photo-carousel";

interface ListingPhoto {
  src: string;
  alt: string;
}

interface ListingMediaSectionProps {
  images: ListingPhoto[];
  description: string;
}

export function ListingMediaSection({
  images,
  description,
}: ListingMediaSectionProps) {
  return (
    <div className="flex h-full flex-col overflow-hidden rounded-[2rem] border border-gray-100 bg-white shadow-xl shadow-orange-900/5">
      <ListingPhotoCarousel images={images} />

      <section className="min-h-[260px] border-t border-gray-100 p-7 sm:p-10">
        <h2 className="mb-4 text-xl font-bold text-gray-800">
          About this space
        </h2>
        <p className="text-base font-medium leading-8 text-gray-500">
          {description}
        </p>
      </section>
    </div>
  );
}
