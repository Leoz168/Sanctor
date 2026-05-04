import Image from "next/image";
import Link from "next/link";
import { ArrowRight, Bath, Bed, MapPin } from "lucide-react";

interface CommunityListingPreviewCardProps {
  href: string;
  image: string;
  price: number;
  title: string;
  location: string;
  beds: number;
  baths: number;
  badge?: string;
}

export function CommunityListingPreviewCard({
  href,
  image,
  price,
  title,
  location,
  beds,
  baths,
  badge,
}: CommunityListingPreviewCardProps) {
  return (
    <Link
      href={href}
      className="group overflow-hidden rounded-[2rem] border border-gray-100 bg-white shadow-sm transition-all hover:-translate-y-1 hover:shadow-xl"
    >
      <div className="relative aspect-[16/10] overflow-hidden">
        <Image
          src={image}
          alt={title}
          fill
          sizes="(min-width: 1024px) 33vw, (min-width: 768px) 50vw, 100vw"
          className="object-cover transition-transform duration-700 group-hover:scale-105"
        />
        {badge && (
          <span className="absolute left-3 top-3 rounded-full bg-brand-orange px-3 py-1 text-[10px] font-black uppercase tracking-[0.14em] text-white shadow-lg shadow-brand-orange/25">
            {badge}
          </span>
        )}
      </div>

      <div className="p-6">
        <div className="mb-2 flex items-start justify-between gap-3">
          <h3 className="text-lg font-bold leading-tight text-gray-900 transition-colors group-hover:text-brand-orange">
            {title}
          </h3>
          <div className="shrink-0 text-right">
            <span className="font-black text-brand-orange">${price}</span>
            <span className="text-sm font-semibold text-gray-400">/mo</span>
          </div>
        </div>

        <div className="mb-5 flex items-center gap-1.5 text-sm font-semibold text-gray-500">
          <MapPin className="h-4 w-4 shrink-0 text-brand-orange/80" />
          <span className="truncate">{location}</span>
        </div>

        <div className="flex items-center justify-between border-t border-gray-100 pt-4">
          <div className="flex items-center gap-4 text-sm font-bold text-gray-500">
            <span className="inline-flex items-center gap-1.5">
              <Bed className="h-4 w-4 text-brand-orange/80" />
              {beds} Bed
            </span>
            <span className="inline-flex items-center gap-1.5">
              <Bath className="h-4 w-4 text-brand-orange/80" />
              {baths} Bath
            </span>
          </div>
          <ArrowRight className="h-5 w-5 text-brand-orange transition-transform group-hover:translate-x-1" />
        </div>
      </div>
    </Link>
  );
}
