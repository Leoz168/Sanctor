"use client";

import Image from "next/image";
import { Star } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";

const reviews = [
  {
    id: 1,
    name: "Nadia Chen",
    date: "February 2026",
    avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Nadia",
    content:
      "The listing felt much more trustworthy because it came through the community. People in the comments answered transit and roommate questions quickly.",
  },
  {
    id: 2,
    name: "Aarav Patel",
    date: "March 2026",
    avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Aarav",
    content:
      "Really appreciated the clear move-in details and the fact that moderators verified the poster before it stayed pinned in the group.",
  },
];

export function Reviews() {
  return (
    <section>
      <div className="mb-6 flex items-center justify-between">
        <h2 className="text-xl font-semibold">Community Feedback</h2>
        <div className="flex items-center gap-2">
          <Star className="h-4 w-4 fill-primary text-primary" />
          <span className="font-semibold">4.9</span>
          <span className="text-sm text-muted-foreground">(38 responses)</span>
        </div>
      </div>

      <div className="space-y-4">
        {reviews.map((review) => (
          <Card key={review.id} className="border">
            <CardContent className="p-4">
              <div className="mb-3 flex items-center gap-3">
                <div className="relative h-10 w-10 overflow-hidden rounded-full bg-muted">
                  <Image
                    src={review.avatar}
                    alt={review.name}
                    fill
                    className="object-cover"
                  />
                </div>
                <div>
                  <p className="text-sm font-medium">{review.name}</p>
                  <p className="text-xs text-muted-foreground">{review.date}</p>
                </div>
              </div>
              <p className="text-sm leading-relaxed text-muted-foreground">
                {review.content}
              </p>
            </CardContent>
          </Card>
        ))}
      </div>

      <Button
        variant="outline"
        className="mt-4 border-primary text-primary hover:bg-primary/5"
      >
        View All Replies
      </Button>
    </section>
  );
}
