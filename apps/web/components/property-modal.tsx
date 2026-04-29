"use client";

import { useState } from "react";
import Image from "next/image";
import { X, Heart, Share2, Radio, Layers3, Users, MapPin, Building, ChevronLeft, ChevronRight } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { Community } from "@/components/property-card";
import { cn } from "@/lib/utils";

interface PropertyModalProps {
  property: Community | null;
  isOpen: boolean;
  onClose: () => void;
}

export function PropertyModal({ property, isOpen, onClose }: PropertyModalProps) {
  const [isFavorited, setIsFavorited] = useState(false);
  const [currentImageIndex, setCurrentImageIndex] = useState(0);

  if (!isOpen || !property) return null;

  const images = [
    property.imageUrl,
    "https://images.unsplash.com/photo-1503676260728-1c00da094a0b?w=800&q=80",
    "https://images.unsplash.com/photo-1529156069898-49953e39b3ac?w=800&q=80",
    "https://images.unsplash.com/photo-1517457373958-b7bdd4587205?w=800&q=80",
  ];

  const nextImage = () => {
    setCurrentImageIndex((prev) => (prev + 1) % images.length);
  };

  const prevImage = () => {
    setCurrentImageIndex((prev) => (prev - 1 + images.length) % images.length);
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
      <div
        className="absolute inset-0 bg-background/80 backdrop-blur-sm"
        onClick={onClose}
      />

      <div className="relative max-h-[90vh] w-full max-w-4xl overflow-hidden rounded-2xl border border-border bg-card shadow-xl">
        <div className="absolute left-0 right-0 top-0 z-10 flex items-center justify-between p-4">
          <button
            onClick={onClose}
            className="flex h-10 w-10 items-center justify-center rounded-full bg-background/80 backdrop-blur-sm transition-colors hover:bg-background"
            aria-label="Close modal"
          >
            <X className="h-5 w-5 text-foreground" />
          </button>
          <div className="flex gap-2">
            <button
              onClick={() => setIsFavorited(!isFavorited)}
              className={cn(
                "flex h-10 w-10 items-center justify-center rounded-full bg-background/80 backdrop-blur-sm transition-colors hover:bg-background",
                isFavorited && "bg-destructive/20 hover:bg-destructive/30",
              )}
              aria-label="Add to favorites"
            >
              <Heart
                className={cn(
                  "h-5 w-5",
                  isFavorited ? "fill-destructive text-destructive" : "text-foreground",
                )}
              />
            </button>
            <button
              className="flex h-10 w-10 items-center justify-center rounded-full bg-background/80 backdrop-blur-sm transition-colors hover:bg-background"
              aria-label="Share community"
            >
              <Share2 className="h-5 w-5 text-foreground" />
            </button>
          </div>
        </div>

        <div className="max-h-[90vh] overflow-y-auto">
          <div className="relative aspect-video bg-secondary">
            <Image
              src={images[currentImageIndex]}
              alt={property.title}
              fill
              className="object-cover"
              sizes="(max-width: 896px) 100vw, 896px"
            />

            <button
              onClick={prevImage}
              className="absolute left-4 top-1/2 flex h-10 w-10 -translate-y-1/2 items-center justify-center rounded-full bg-background/80 backdrop-blur-sm transition-colors hover:bg-background"
              aria-label="Previous image"
            >
              <ChevronLeft className="h-5 w-5 text-foreground" />
            </button>
            <button
              onClick={nextImage}
              className="absolute right-4 top-1/2 flex h-10 w-10 -translate-y-1/2 items-center justify-center rounded-full bg-background/80 backdrop-blur-sm transition-colors hover:bg-background"
              aria-label="Next image"
            >
              <ChevronRight className="h-5 w-5 text-foreground" />
            </button>

            <div className="absolute bottom-4 left-1/2 flex -translate-x-1/2 gap-2">
              {images.map((_, index) => (
                <button
                  key={index}
                  onClick={() => setCurrentImageIndex(index)}
                  className={cn(
                    "h-2 w-2 rounded-full transition-colors",
                    index === currentImageIndex ? "bg-foreground" : "bg-foreground/40",
                  )}
                  aria-label={`Go to image ${index + 1}`}
                />
              ))}
            </div>
          </div>

          <div className="p-6">
            {property.status && (
              <Badge
                className={cn(
                  "mb-3",
                  property.status === "new" && "bg-accent text-accent-foreground",
                  property.status === "featured" && "bg-primary text-primary-foreground",
                  property.status === "reduced" && "bg-destructive text-destructive-foreground",
                )}
              >
                {property.status === "new" && "New Community"}
                {property.status === "featured" && "Featured"}
                {property.status === "reduced" && "Selective Access"}
              </Badge>
            )}

            <div className="mb-4">
              <h2 className="text-2xl font-bold text-foreground">
                {property.members.toLocaleString()} members
              </h2>
              <h3 className="mt-1 text-lg text-foreground">{property.title}</h3>
              <div className="mt-2 flex items-center gap-1 text-muted-foreground">
                <MapPin className="h-4 w-4" />
                <span>{property.institution} • {property.location}</span>
              </div>
            </div>

            <div className="mb-6 flex flex-wrap gap-6">
              <div className="flex items-center gap-2">
                <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-secondary">
                  <Users className="h-5 w-5 text-muted-foreground" />
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Members</p>
                  <p className="font-medium text-foreground">{property.members.toLocaleString()}</p>
                </div>
              </div>
              <div className="flex items-center gap-2">
                <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-secondary">
                  <Radio className="h-5 w-5 text-muted-foreground" />
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Online Now</p>
                  <p className="font-medium text-foreground">{property.onlineNow}</p>
                </div>
              </div>
              <div className="flex items-center gap-2">
                <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-secondary">
                  <Layers3 className="h-5 w-5 text-muted-foreground" />
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Channels</p>
                  <p className="font-medium text-foreground">{property.channels}</p>
                </div>
              </div>
              <div className="flex items-center gap-2">
                <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-secondary">
                  <Building className="h-5 w-5 text-muted-foreground" />
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Community Type</p>
                  <p className="font-medium capitalize text-foreground">{property.type}</p>
                </div>
              </div>
            </div>

            <Separator className="my-6" />

            <div className="mb-6">
              <h4 className="mb-3 text-lg font-semibold text-foreground">About This Community</h4>
              <p className="leading-relaxed text-muted-foreground">
                {property.description}
              </p>
            </div>

            <div className="mb-6">
              <h4 className="mb-3 text-lg font-semibold text-foreground">Focus Areas</h4>
              <div className="grid grid-cols-2 gap-2 sm:grid-cols-3">
                {property.tags.map((feature) => (
                  <div key={feature} className="flex items-center gap-2 text-sm text-muted-foreground">
                    <div className="h-1.5 w-1.5 rounded-full bg-accent" />
                    {feature}
                  </div>
                ))}
              </div>
            </div>

            <Separator className="my-6" />

            <div className="flex flex-col gap-4 sm:flex-row">
              <Button className="flex-1 bg-primary text-primary-foreground hover:bg-primary/90">
                Join Community
              </Button>
              <Button variant="outline" className="flex-1 border-border text-foreground hover:bg-secondary">
                Message Moderators
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
