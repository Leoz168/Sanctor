import Image from "next/image";
import { ArrowRight, Bath, Bed, Bookmark, MapPin } from "lucide-react";
import { Button } from "@/components/ui/button";

interface ListingCardProps {
  title: string;
  price: number;
  location: string;
  beds: number;
  baths: number;
  image: string;
  badge?: "featured" | "new";
}

export function ListingCard({
  title,
  price,
  location,
  beds,
  baths,
  image,
  badge,
}: ListingCardProps) {
  return (
    <div className="bg-white rounded-[2rem] border border-gray-100 overflow-hidden group shadow-sm hover:shadow-xl hover:-translate-y-1 transition-all">
      <div className="relative aspect-[16/10] overflow-hidden">
        <Image
          src={image}
          alt={title}
          fill
          sizes="(min-width: 1024px) 33vw, (min-width: 768px) 50vw, 100vw"
          className="object-cover transition-transform duration-700 group-hover:scale-105"
        />
        {badge && (
          <span className="absolute top-3 left-3 px-3 py-1 text-[10px] font-bold uppercase tracking-wider bg-brand-orange text-white rounded-full shadow-lg">
            {badge}
          </span>
        )}
        <Button
          size="icon"
          variant="outline"
          className="absolute top-3 right-3 w-9 h-9 rounded-full bg-white/85 backdrop-blur-sm border-white/70 text-gray-900 hover:bg-white hover:text-brand-orange"
          aria-label={`Save listing: ${title}`}
          title="Save listing"
        >
          <Bookmark className="w-4 h-4" />
        </Button>
      </div>

      <div className="p-6">
        <div className="flex items-start justify-between gap-3 mb-2">
          <h3 className="font-bold text-lg text-gray-900 leading-tight group-hover:text-brand-orange transition-colors">
            {title}
          </h3>
          <div className="text-right shrink-0">
            <span className="text-brand-orange font-bold">${price}</span>
            <span className="text-gray-400 text-sm">/mo</span>
          </div>
        </div>

        <div className="flex items-center gap-1.5 text-sm text-gray-500 mb-5 font-medium">
          <MapPin className="w-4 h-4 shrink-0" />
          <span className="truncate">{location}</span>
        </div>

        <div className="flex items-center justify-between pt-4 border-t border-gray-100">
          <div className="flex items-center gap-4 text-sm text-gray-500 font-semibold">
            <div className="flex items-center gap-1.5">
              <Bed className="w-4 h-4 text-brand-orange/70" />
              <span>{beds} Bed</span>
            </div>
            <div className="flex items-center gap-1.5">
              <Bath className="w-4 h-4 text-brand-orange/70" />
              <span>{baths} Bath</span>
            </div>
          </div>
          <Button
            size="icon"
            variant="ghost"
            className="w-8 h-8 text-brand-orange hover:text-brand-orange hover:bg-brand-orange/10"
            aria-label={`View ${title}`}
          >
            <ArrowRight className="w-4 h-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
