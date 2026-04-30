"use client";

import Image from "next/image";

const images = [
  { src: "https://images.unsplash.com/photo-1502672260266-1c1ef2d93688?w=1600&q=80", alt: "Shared living room with sunlight" },
  { src: "https://images.unsplash.com/photo-1484154218962-a197022b5858?w=1200&q=80", alt: "Kitchen and dining area" },
  { src: "https://images.unsplash.com/photo-1505693416388-ac5ce068fe85?w=1200&q=80", alt: "Bedroom space" },
  { src: "https://images.unsplash.com/photo-1513694203232-719a280e022f?w=1200&q=80", alt: "Study corner near the window" },
];

export function ImageGallery() {
  return (
    <div className="grid grid-cols-1 gap-2 md:h-[520px] md:grid-cols-3 md:grid-rows-2">
      <div className="relative aspect-[4/3] overflow-hidden rounded-lg md:col-span-2 md:row-span-2 md:aspect-auto md:h-full">
        <span className="absolute left-3 top-3 z-10 rounded bg-primary px-3 py-1 text-xs font-semibold text-primary-foreground">
          COMMUNITY LISTING
        </span>
        <Image
          src={images[0].src}
          alt={images[0].alt}
          fill
          className="object-cover"
          priority
        />
      </div>

      {images.slice(1).map((image, index) => (
        <div
          key={image.alt}
          className="relative hidden overflow-hidden rounded-lg md:block"
        >
          <Image
            src={image.src}
            alt={image.alt}
            fill
            className="object-cover"
          />
          {index === 2 ? (
            <div className="absolute inset-0 flex items-center justify-center bg-black/50">
              <span className="text-xl font-semibold text-white">+8</span>
            </div>
          ) : null}
        </div>
      ))}
    </div>
  );
}
