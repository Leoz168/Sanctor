import { Amenities } from "@/components/listing/amenities";
import { ListingFooter } from "@/components/listing/listing-footer";
import { ListingHeader } from "@/components/listing/listing-header";
import { ImageGallery } from "@/components/listing/image-gallery";
import { ListingDetails } from "@/components/listing/listing-details";
import { PremiumBanner } from "@/components/listing/premium-banner";
import { PricingCard } from "@/components/listing/pricing-card";
import { Reviews } from "@/components/listing/reviews";

export default function CommunityListingPage() {
  return (
    <div className="flex min-h-screen flex-col bg-background">
      <ListingHeader />

      <main className="flex-1">
        <div className="container mx-auto px-4 py-8">
          <div className="grid grid-cols-1 gap-8 lg:grid-cols-3">
            <div className="space-y-8 lg:col-span-2">
              <ImageGallery />
              <ListingDetails />
              <hr className="border-border" />
              <Amenities />
              <hr className="border-border" />
              <Reviews />
            </div>

            <div className="space-y-6">
              <PricingCard />
              <PremiumBanner />
            </div>
          </div>
        </div>
      </main>

      <ListingFooter />
    </div>
  );
}
