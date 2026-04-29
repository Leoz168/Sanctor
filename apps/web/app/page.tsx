"use client";

import { useState } from "react";
import { ChevronDown, MapPin, SlidersHorizontal } from "lucide-react";
import { Header } from "@/components/header";
import { SearchBar, SearchFilters } from "@/components/search-bar";
import { PropertyGrid } from "@/components/property-grid";
import { FilterSidebar, Filters } from "@/components/filter-sidebar";
import { PropertyModal } from "@/components/property-modal";
import { Community } from "@/components/property-card";
import { Button } from "@/components/ui/button";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const sampleCommunities: Community[] = [
  {
    id: "1",
    title: "Computer Science Collective",
    institution: "University of Toronto",
    location: "St. George Campus",
    members: 1248,
    onlineNow: 216,
    weeklyPosts: 382,
    channels: 18,
    imageUrl: "https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=1200&q=80",
    type: "academic",
    status: "featured",
    description: "Project nights, interview prep, hackathon teams, and daily student support for developers across campus.",
    tags: ["Hackathons", "Study Groups", "Career Prep", "Peer Help"],
  },
  {
    id: "2",
    title: "Design and Product Guild",
    institution: "Toronto Metropolitan University",
    location: "Downtown Toronto",
    members: 624,
    onlineNow: 84,
    weeklyPosts: 147,
    channels: 11,
    imageUrl: "https://images.unsplash.com/photo-1516321318423-f06f85e504b3?w=1200&q=80",
    type: "professional",
    status: "new",
    description: "A home for product thinkers, designers, and founders sharing critiques, workshops, and portfolio feedback.",
    tags: ["Portfolio Reviews", "Workshops", "Startups"],
  },
  {
    id: "3",
    title: "Residence Life Hub",
    institution: "McMaster University",
    location: "Hamilton",
    members: 1892,
    onlineNow: 331,
    weeklyPosts: 524,
    channels: 24,
    imageUrl: "https://images.unsplash.com/photo-1523240795612-9a054b0db644?w=1200&q=80",
    type: "social",
    description: "Meet neighbors, trade advice, coordinate events, and keep your residence community connected every day.",
    tags: ["Events", "Housing", "Support"],
  },
  {
    id: "4",
    title: "Pre-Med Peer Network",
    institution: "Western University",
    location: "London, Ontario",
    members: 812,
    onlineNow: 129,
    weeklyPosts: 201,
    channels: 9,
    imageUrl: "https://images.unsplash.com/photo-1571260899304-425eee4c7efc?w=1200&q=80",
    type: "academic",
    status: "reduced",
    description: "Course notes, MCAT accountability, mentorship, and application support for future healthcare professionals.",
    tags: ["Mentorship", "MCAT", "Course Help"],
  },
  {
    id: "5",
    title: "Campus Creators Club",
    institution: "York University",
    location: "Keele Campus",
    members: 458,
    onlineNow: 61,
    weeklyPosts: 96,
    channels: 7,
    imageUrl: "https://images.unsplash.com/photo-1516321497487-e288fb19713f?w=1200&q=80",
    type: "creator",
    description: "For filmmakers, writers, streamers, and visual storytellers building ambitious work together.",
    tags: ["Film", "Writing", "Content Creation"],
  },
  {
    id: "6",
    title: "International Student Circle",
    institution: "University of Waterloo",
    location: "Waterloo",
    members: 1405,
    onlineNow: 278,
    weeklyPosts: 430,
    channels: 16,
    imageUrl: "https://images.unsplash.com/photo-1529390079861-591de354faf5?w=1200&q=80",
    type: "support",
    status: "featured",
    description: "Arrival tips, immigration Q&A, friendship-building, and real-time help for students adjusting to campus life.",
    tags: ["Arrival", "Immigration", "Community"],
  },
];

export default function HomePage() {
  const [filterSidebarOpen, setFilterSidebarOpen] = useState(false);
  const [selectedCommunity, setSelectedCommunity] = useState<Community | null>(null);
  const [sortBy, setSortBy] = useState("activity");
  const [communities] = useState(sampleCommunities);

  const handleSearch = (filters: SearchFilters) => {
    console.log("Search filters:", filters);
  };

  const handleApplyFilters = (filters: Filters) => {
    console.log("Applied filters:", filters);
  };

  const handleCommunityClick = (id: string) => {
    const community = communities.find((item) => item.id === id);
    if (community) {
      setSelectedCommunity(community);
    }
  };

  const handleFavorite = (id: string) => {
    console.log("Favorited community:", id);
  };

  return (
    <div className="min-h-screen bg-background">
      <Header />

      <section className="border-b border-border bg-card px-4 py-12 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-7xl">
          <div className="mb-8 text-center">
            <h1 className="text-balance text-3xl font-bold tracking-tight text-foreground sm:text-4xl lg:text-5xl">
              Discover campus communities that actually feel alive.
            </h1>
            <p className="mx-auto mt-4 max-w-2xl text-pretty text-muted-foreground">
              Browse student groups, residence hubs, academic circles, and support networks across your institution with a sharper
              Sanctor experience.
            </p>
          </div>
          <SearchBar onSearch={handleSearch} />
        </div>
      </section>

      <main className="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
        <div className="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div className="flex items-center gap-2">
            <MapPin className="h-4 w-4 text-muted-foreground" />
            <span className="text-sm text-muted-foreground">
              <span className="font-medium text-foreground">{communities.length}</span> active communities across Ontario campuses
            </span>
          </div>

          <div className="flex items-center gap-3">
            <Select value={sortBy} onValueChange={setSortBy}>
              <SelectTrigger className="w-[180px] border-border bg-card text-foreground">
                <SelectValue placeholder="Sort by" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="activity">Most Active</SelectItem>
                <SelectItem value="members">Largest Communities</SelectItem>
                <SelectItem value="newest">Newest Groups</SelectItem>
                <SelectItem value="online">Most Online Now</SelectItem>
              </SelectContent>
            </Select>

            <Button
              variant="outline"
              className="gap-2 border-border text-foreground hover:bg-secondary"
              onClick={() => setFilterSidebarOpen(true)}
            >
              <SlidersHorizontal className="h-4 w-4" />
              <span className="hidden sm:inline">Filters</span>
            </Button>
          </div>
        </div>

        <PropertyGrid
          properties={communities}
          onPropertyClick={handleCommunityClick}
          onFavorite={handleFavorite}
        />

        <div className="mt-10 flex justify-center">
          <Button
            variant="outline"
            size="lg"
            className="gap-2 border-border text-foreground hover:bg-secondary"
          >
            Load More Communities
            <ChevronDown className="h-4 w-4" />
          </Button>
        </div>
      </main>

      <footer className="border-t border-border bg-card">
        <div className="mx-auto max-w-7xl px-4 py-12 sm:px-6 lg:px-8">
          <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-4">
            <div>
              <h3 className="text-sm font-semibold text-foreground">Platform</h3>
              <ul className="mt-4 space-y-2">
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Communities</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Messaging</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Moderation</a></li>
              </ul>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-foreground">Support</h3>
              <ul className="mt-4 space-y-2">
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Help Center</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Community Safety</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Contact Us</a></li>
              </ul>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-foreground">Developers</h3>
              <ul className="mt-4 space-y-2">
                <li><a href={process.env.NEXT_PUBLIC_API_URL ?? "#"} className="text-sm text-muted-foreground hover:text-foreground">API</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Integrations</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Status</a></li>
              </ul>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-foreground">Connect</h3>
              <ul className="mt-4 space-y-2">
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Instagram</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">LinkedIn</a></li>
                <li><a href="#" className="text-sm text-muted-foreground hover:text-foreground">Campus Partners</a></li>
              </ul>
            </div>
          </div>
          <div className="mt-8 border-t border-border pt-8 text-center">
            <p className="text-sm text-muted-foreground">
              &copy; {new Date().getFullYear()} Sanctor. Community infrastructure for modern campuses.
            </p>
          </div>
        </div>
      </footer>

      <FilterSidebar
        isOpen={filterSidebarOpen}
        onClose={() => setFilterSidebarOpen(false)}
        onApplyFilters={handleApplyFilters}
      />

      <PropertyModal
        property={selectedCommunity}
        isOpen={!!selectedCommunity}
        onClose={() => setSelectedCommunity(null)}
      />
    </div>
  );
}
