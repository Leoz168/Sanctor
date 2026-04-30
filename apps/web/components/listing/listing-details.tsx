"use client";

import { Heart, MapPin, Share2, Users } from "lucide-react";
import { Button } from "@/components/ui/button";

export function ListingDetails() {
  return (
    <div>
      <div className="mb-2 flex items-start justify-between gap-4">
        <div>
          <p className="mb-2 text-sm font-medium text-accent">U of T Off-Campus Housing • Housing Listing</p>
          <h1 className="text-balance text-2xl font-bold md:text-3xl">
            Bright Room in Harbord Village Townhouse
          </h1>
        </div>
        <div className="flex shrink-0 gap-2">
          <Button variant="outline" size="icon" className="h-9 w-9">
            <Share2 className="h-4 w-4" />
          </Button>
          <Button variant="outline" size="icon" className="h-9 w-9">
            <Heart className="h-4 w-4" />
          </Button>
        </div>
      </div>

      <div className="mb-3 flex flex-wrap items-center gap-4 text-muted-foreground">
        <div className="flex items-center gap-2 text-sm">
          <MapPin className="h-4 w-4" />
          <span>Harbord Village, Toronto</span>
        </div>
        <div className="flex items-center gap-2 text-sm">
          <Users className="h-4 w-4" />
          <span>Posted inside a 1.4k-member housing community</span>
        </div>
      </div>

      <p className="leading-relaxed text-muted-foreground">
        This listing was shared in Sanctor by a student moderator-approved member of the U of T Off-Campus Housing community.
        The room is available for a summer sublet with flexible move-in dates, a furnished common area, and fast TTC access.
        It is designed to feel trustworthy, social, and clearly connected to the community it belongs to rather than appearing as a detached marketplace post.
      </p>
    </div>
  );
}
