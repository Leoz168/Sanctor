"use client";

import { PropertyCard, Community } from "@/components/property-card";

interface PropertyGridProps {
  properties: Community[];
  onPropertyClick?: (id: string) => void;
  onFavorite?: (id: string) => void;
}

export function PropertyGrid({ properties, onPropertyClick, onFavorite }: PropertyGridProps) {
  if (properties.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center py-16 text-center">
        <div className="mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-secondary">
          <svg
            className="h-8 w-8 text-muted-foreground"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={1.5}
              d="M4 8h16M4 12h16M4 16h10"
            />
          </svg>
        </div>
        <h3 className="text-lg font-medium text-foreground">No communities found</h3>
        <p className="mt-1 text-sm text-muted-foreground">
          Try adjusting your filters or search criteria
        </p>
      </div>
    );
  }

  return (
    <div className="grid gap-6 sm:grid-cols-2 xl:grid-cols-3">
      {properties.map((property) => (
        <PropertyCard
          key={property.id}
          property={property}
          onClick={onPropertyClick}
          onFavorite={onFavorite}
        />
      ))}
    </div>
  );
}
