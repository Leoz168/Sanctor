"use client";

import { useState } from "react";
import Image from "next/image";
import { Heart, Radio, Layers3, ShieldCheck, MapPin } from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";

export interface Community {
  id: string;
  title: string;
  institution: string;
  location: string;
  members: number;
  onlineNow: number;
  weeklyPosts: number;
  channels: number;
  imageUrl: string;
  type: "academic" | "social" | "support" | "creator" | "professional";
  status?: "new" | "featured" | "reduced";
  description: string;
  tags: string[];
}

interface PropertyCardProps {
  property: Community;
  onFavorite?: (id: string) => void;
  onClick?: (id: string) => void;
}

export function PropertyCard({ property, onFavorite, onClick }: PropertyCardProps) {
  const [isFavorited, setIsFavorited] = useState(false);
  const [imageLoaded, setImageLoaded] = useState(false);

  const handleFavorite = (e: React.MouseEvent) => {
    e.stopPropagation();
    setIsFavorited(!isFavorited);
    onFavorite?.(property.id);
  };

  return (
    <article
      onClick={() => onClick?.(property.id)}
      className="group cursor-pointer overflow-hidden rounded-xl border border-border bg-card transition-all hover:border-accent/50 hover:shadow-lg hover:shadow-accent/5"
    >
      <div className="relative aspect-[4/3] overflow-hidden bg-secondary">
        <Image
          src={property.imageUrl}
          alt={property.title}
          fill
          className={cn(
            "object-cover transition-all duration-300 group-hover:scale-105",
            imageLoaded ? "opacity-100" : "opacity-0",
          )}
          onLoad={() => setImageLoaded(true)}
          sizes="(max-width: 640px) 100vw, (max-width: 1024px) 50vw, 33vw"
        />
        {!imageLoaded && (
          <div className="absolute inset-0 animate-pulse bg-secondary" />
        )}

        {property.status && (
          <Badge
            className={cn(
              "absolute left-3 top-3 font-medium",
              property.status === "new" && "bg-accent text-accent-foreground",
              property.status === "featured" && "bg-primary text-primary-foreground",
              property.status === "reduced" && "bg-destructive text-destructive-foreground",
            )}
          >
            {property.status === "new" && "New"}
            {property.status === "featured" && "Featured"}
            {property.status === "reduced" && "Selective Access"}
          </Badge>
        )}

        <button
          onClick={handleFavorite}
          className={cn(
            "absolute right-3 top-3 flex h-9 w-9 items-center justify-center rounded-full bg-background/80 backdrop-blur-sm transition-all hover:bg-background",
            isFavorited && "bg-destructive/20 hover:bg-destructive/30",
          )}
          aria-label={isFavorited ? "Remove from favorites" : "Add to favorites"}
        >
          <Heart
            className={cn(
              "h-5 w-5 transition-colors",
              isFavorited ? "fill-destructive text-destructive" : "text-foreground",
            )}
          />
        </button>

        <div className="absolute bottom-3 left-3">
          <Badge variant="secondary" className="bg-background/80 text-foreground backdrop-blur-sm">
            {property.type}
          </Badge>
        </div>
      </div>

      <div className="p-4">
        <div className="mb-2 flex items-baseline justify-between">
          <span className="text-xl font-bold text-foreground">
            {property.members.toLocaleString()} members
          </span>
          <span className="text-sm text-muted-foreground">{property.weeklyPosts} posts this week</span>
        </div>

        <h3 className="mb-1 line-clamp-1 text-base font-medium text-foreground">
          {property.title}
        </h3>

        <div className="mb-3 flex items-center gap-1 text-sm text-muted-foreground">
          <MapPin className="h-3.5 w-3.5 flex-shrink-0" />
          <span className="line-clamp-1">{property.institution} • {property.location}</span>
        </div>

        <div className="flex items-center gap-4 text-sm text-muted-foreground">
          <div className="flex items-center gap-1.5">
            <Radio className="h-4 w-4" />
            <span>{property.onlineNow} online</span>
          </div>
          <div className="flex items-center gap-1.5">
            <Layers3 className="h-4 w-4" />
            <span>{property.channels} channels</span>
          </div>
          <div className="flex items-center gap-1.5">
            <ShieldCheck className="h-4 w-4" />
            <span>{property.tags.length} focus areas</span>
          </div>
        </div>
      </div>
    </article>
  );
}
