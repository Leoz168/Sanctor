import Image from "next/image";
import Link from "next/link";
import { ArrowRight, Bath, BedDouble, MapPin } from "lucide-react";

interface PropertyCardProps {
  href: string;
  image: string;
  price: string;
  title: string;
  location: string;
  beds: number;
  baths: number;
}

export function PropertyCard({
  href,
  image,
  price,
  title,
  location,
  beds,
  baths,
}: PropertyCardProps) {
  return (
    <Link href={href} className="group block cursor-pointer">
      <div className="relative mb-4 aspect-[4/3] overflow-hidden rounded-xl">
        <Image
          src={image}
          alt={title}
          fill
          className="object-cover transition-transform duration-300 group-hover:scale-105"
        />
        <div className="absolute right-3 top-3">
          <span className="rounded-full bg-primary px-3 py-1 text-sm font-semibold text-primary-foreground">
            {price}
          </span>
        </div>
      </div>

      <div>
        <h3 className="mb-1 font-semibold text-foreground">{title}</h3>
        <div className="mb-3 flex items-center gap-1 text-sm text-muted-foreground">
          <MapPin className="h-4 w-4" />
          <span>{location}</span>
        </div>

        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4 text-sm text-muted-foreground">
            <div className="flex items-center gap-1">
              <BedDouble className="h-4 w-4" />
              <span>{beds} Bed</span>
            </div>
            <div className="flex items-center gap-1">
              <Bath className="h-4 w-4" />
              <span>{baths} Bath</span>
            </div>
          </div>
          <ArrowRight className="h-5 w-5 text-primary transition-transform group-hover:translate-x-1" />
        </div>
      </div>
    </Link>
  );
}
