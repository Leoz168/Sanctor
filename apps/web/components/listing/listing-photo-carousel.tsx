"use client";

import { useState } from "react";
import Image from "next/image";
import { ChevronLeft, ChevronRight } from "lucide-react";

interface ListingPhoto {
  src: string;
  alt: string;
}

interface ListingPhotoCarouselProps {
  images: ListingPhoto[];
}

export function ListingPhotoCarousel({ images }: ListingPhotoCarouselProps) {
  const [currentIndex, setCurrentIndex] = useState(0);
  const currentImage = images[currentIndex];

  const showPrevious = () => {
    setCurrentIndex((index) => (index === 0 ? images.length - 1 : index - 1));
  };

  const showNext = () => {
    setCurrentIndex((index) => (index === images.length - 1 ? 0 : index + 1));
  };

  return (
    <section className="relative min-h-[290px] flex-1 overflow-hidden sm:min-h-[375px] lg:min-h-[430px]">
      <Image
        src={currentImage.src}
        alt={currentImage.alt}
        fill
        priority
        sizes="(min-width: 1024px) 68vw, 100vw"
        className="object-cover"
      />

      <div className="absolute left-5 top-5 rounded-full bg-brand-orange px-4 py-2 text-xs font-black uppercase tracking-[0.14em] text-white shadow-lg shadow-brand-orange/25">
        Verified listing
      </div>

      <button
        type="button"
        onClick={showPrevious}
        className="absolute left-4 top-1/2 inline-flex h-11 w-11 -translate-y-1/2 items-center justify-center rounded-full bg-white/90 text-gray-800 shadow-lg backdrop-blur transition-colors hover:text-brand-orange"
        aria-label="Show previous photo"
      >
        <ChevronLeft size={24} />
      </button>

      <button
        type="button"
        onClick={showNext}
        className="absolute right-4 top-1/2 inline-flex h-11 w-11 -translate-y-1/2 items-center justify-center rounded-full bg-white/90 text-gray-800 shadow-lg backdrop-blur transition-colors hover:text-brand-orange"
        aria-label="Show next photo"
      >
        <ChevronRight size={24} />
      </button>

      <div className="absolute bottom-5 left-1/2 flex -translate-x-1/2 gap-2 rounded-full bg-white/85 px-3 py-2 shadow-lg backdrop-blur">
        {images.map((image, index) => (
          <button
            key={image.src}
            type="button"
            onClick={() => setCurrentIndex(index)}
            className={`h-2.5 rounded-full transition-all ${
              index === currentIndex
                ? "w-7 bg-brand-orange"
                : "w-2.5 bg-gray-300 hover:bg-gray-400"
            }`}
            aria-label={`Show photo ${index + 1}`}
          />
        ))}
      </div>
    </section>
  );
}
