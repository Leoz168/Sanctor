"use client";

import { useState } from "react";
import { Search, MapPin } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface SearchBarProps {
  onSearch?: (filters: SearchFilters) => void;
}

export interface SearchFilters {
  institution: string;
  communityType: string;
  communitySize: string;
  accessType: string;
}

export function SearchBar({ onSearch }: SearchBarProps) {
  const [institution, setInstitution] = useState("");
  const [communityType, setCommunityType] = useState("");
  const [communitySize, setCommunitySize] = useState("");
  const [accessType, setAccessType] = useState("");

  const handleSearch = () => {
    onSearch?.({
      institution,
      communityType,
      communitySize,
      accessType,
    });
  };

  return (
    <div className="w-full rounded-xl border border-border bg-card p-4">
      <div className="flex flex-col gap-4 lg:flex-row lg:items-end">
        <div className="flex-1">
          <label className="mb-2 block text-sm font-medium text-muted-foreground">
            Institution
          </label>
          <div className="relative">
            <MapPin className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="School, campus, or city"
              value={institution}
              onChange={(e) => setInstitution(e.target.value)}
              className="h-11 border-border bg-secondary pl-10 text-foreground placeholder:text-muted-foreground"
            />
          </div>
        </div>

        <div className="flex-1 lg:max-w-[180px]">
          <label className="mb-2 block text-sm font-medium text-muted-foreground">
            Community Type
          </label>
          <Select value={communityType} onValueChange={setCommunityType}>
            <SelectTrigger className="h-11 border-border bg-secondary text-foreground">
              <SelectValue placeholder="Any type" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="academic">Academic</SelectItem>
              <SelectItem value="social">Social</SelectItem>
              <SelectItem value="support">Support</SelectItem>
              <SelectItem value="creator">Creator</SelectItem>
              <SelectItem value="professional">Professional</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <div className="flex-1 lg:max-w-[180px]">
          <label className="mb-2 block text-sm font-medium text-muted-foreground">
            Community Size
          </label>
          <Select value={communitySize} onValueChange={setCommunitySize}>
            <SelectTrigger className="h-11 border-border bg-secondary text-foreground">
              <SelectValue placeholder="Any size" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="0-100">Under 100</SelectItem>
              <SelectItem value="100-500">100 - 500</SelectItem>
              <SelectItem value="500-1000">500 - 1,000</SelectItem>
              <SelectItem value="1000-2000">1,000 - 2,000</SelectItem>
              <SelectItem value="2000+">2,000+</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <div className="flex-1 lg:max-w-[140px]">
          <label className="mb-2 block text-sm font-medium text-muted-foreground">
            Access
          </label>
          <Select value={accessType} onValueChange={setAccessType}>
            <SelectTrigger className="h-11 border-border bg-secondary text-foreground">
              <SelectValue placeholder="Any" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="open">Open</SelectItem>
              <SelectItem value="request">Request Access</SelectItem>
              <SelectItem value="verified">Verified Only</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <Button
          onClick={handleSearch}
          className="h-11 gap-2 bg-primary px-6 text-primary-foreground hover:bg-primary/90"
        >
          <Search className="h-4 w-4" />
          Search
        </Button>
      </div>
    </div>
  );
}
