"use client";

import { useState } from "react";
import { X, ChevronDown, ChevronUp } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { Slider } from "@/components/ui/slider";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";

interface FilterSidebarProps {
  isOpen: boolean;
  onClose: () => void;
  onApplyFilters?: (filters: Filters) => void;
}

export interface Filters {
  memberRange: [number, number];
  communityTypes: string[];
  activityLevel: number | null;
  moderationLevel: number | null;
  features: string[];
}

const communityTypes = [
  { id: "academic", label: "Academic" },
  { id: "social", label: "Social" },
  { id: "support", label: "Support" },
  { id: "creator", label: "Creator" },
  { id: "professional", label: "Professional" },
];

const features = [
  { id: "verified", label: "Verified Access" },
  { id: "events", label: "Live Events" },
  { id: "mentorship", label: "Mentorship" },
  { id: "resources", label: "Resource Library" },
  { id: "roles", label: "Role Channels" },
  { id: "dm", label: "Direct Messaging" },
  { id: "mod", label: "Active Moderation" },
  { id: "bots", label: "Bot Integrations" },
];

export function FilterSidebar({ isOpen, onClose, onApplyFilters }: FilterSidebarProps) {
  const [memberRange, setMemberRange] = useState<[number, number]>([0, 5000]);
  const [selectedCommunityTypes, setSelectedCommunityTypes] = useState<string[]>([]);
  const [activityLevel, setActivityLevel] = useState<number | null>(null);
  const [moderationLevel, setModerationLevel] = useState<number | null>(null);
  const [selectedFeatures, setSelectedFeatures] = useState<string[]>([]);
  const [expandedSections, setExpandedSections] = useState<string[]>(["members", "type", "activity", "features"]);

  const toggleSection = (section: string) => {
    setExpandedSections((prev) =>
      prev.includes(section) ? prev.filter((item) => item !== section) : [...prev, section],
    );
  };

  const toggleCommunityType = (type: string) => {
    setSelectedCommunityTypes((prev) =>
      prev.includes(type) ? prev.filter((item) => item !== type) : [...prev, type],
    );
  };

  const toggleFeature = (feature: string) => {
    setSelectedFeatures((prev) =>
      prev.includes(feature) ? prev.filter((item) => item !== feature) : [...prev, feature],
    );
  };

  const handleApply = () => {
    onApplyFilters?.({
      memberRange,
      communityTypes: selectedCommunityTypes,
      activityLevel,
      moderationLevel,
      features: selectedFeatures,
    });
    onClose();
  };

  const handleReset = () => {
    setMemberRange([0, 5000]);
    setSelectedCommunityTypes([]);
    setActivityLevel(null);
    setModerationLevel(null);
    setSelectedFeatures([]);
  };

  const formatMembers = (value: number) => {
    if (value >= 1000) {
      return `${(value / 1000).toFixed(1)}k`;
    }
    return `${value}`;
  };

  return (
    <>
      {isOpen && (
        <div
          className="fixed inset-0 z-40 bg-background/80 backdrop-blur-sm lg:hidden"
          onClick={onClose}
        />
      )}

      <aside
        className={`fixed right-0 top-0 z-50 h-full w-full max-w-sm transform overflow-y-auto border-l border-border bg-background transition-transform duration-300 lg:static lg:z-0 lg:block lg:max-w-none lg:transform-none lg:border-l-0 lg:border-r ${
          isOpen ? "translate-x-0" : "translate-x-full lg:translate-x-0"
        }`}
      >
        <div className="flex items-center justify-between border-b border-border p-4 lg:hidden">
          <h2 className="text-lg font-semibold text-foreground">Filters</h2>
          <button onClick={onClose} aria-label="Close filters">
            <X className="h-5 w-5 text-muted-foreground" />
          </button>
        </div>

        <div className="p-4">
          <div className="mb-4">
            <button
              onClick={() => toggleSection("members")}
              className="flex w-full items-center justify-between py-2"
            >
              <span className="text-sm font-medium text-foreground">Member Count</span>
              {expandedSections.includes("members") ? (
                <ChevronUp className="h-4 w-4 text-muted-foreground" />
              ) : (
                <ChevronDown className="h-4 w-4 text-muted-foreground" />
              )}
            </button>
            {expandedSections.includes("members") && (
              <div className="mt-4 space-y-4">
                <Slider
                  min={0}
                  max={5000}
                  step={50}
                  value={memberRange}
                  onValueChange={(value: number[]) => setMemberRange(value as [number, number])}
                  className="w-full"
                />
                <div className="flex items-center justify-between text-sm text-muted-foreground">
                  <span>{formatMembers(memberRange[0])}</span>
                  <span>{formatMembers(memberRange[1])}</span>
                </div>
              </div>
            )}
          </div>

          <Separator className="my-4" />

          <div className="mb-4">
            <button
              onClick={() => toggleSection("type")}
              className="flex w-full items-center justify-between py-2"
            >
              <span className="text-sm font-medium text-foreground">Community Type</span>
              {expandedSections.includes("type") ? (
                <ChevronUp className="h-4 w-4 text-muted-foreground" />
              ) : (
                <ChevronDown className="h-4 w-4 text-muted-foreground" />
              )}
            </button>
            {expandedSections.includes("type") && (
              <div className="mt-3 space-y-3">
                {communityTypes.map((type) => (
                  <div key={type.id} className="flex items-center gap-2">
                    <Checkbox
                      id={type.id}
                      checked={selectedCommunityTypes.includes(type.id)}
                      onCheckedChange={() => toggleCommunityType(type.id)}
                    />
                    <Label htmlFor={type.id} className="cursor-pointer text-sm text-muted-foreground">
                      {type.label}
                    </Label>
                  </div>
                ))}
              </div>
            )}
          </div>

          <Separator className="my-4" />

          <div className="mb-4">
            <button
              onClick={() => toggleSection("activity")}
              className="flex w-full items-center justify-between py-2"
            >
              <span className="text-sm font-medium text-foreground">Activity & Moderation</span>
              {expandedSections.includes("activity") ? (
                <ChevronUp className="h-4 w-4 text-muted-foreground" />
              ) : (
                <ChevronDown className="h-4 w-4 text-muted-foreground" />
              )}
            </button>
            {expandedSections.includes("activity") && (
              <div className="mt-3 space-y-4">
                <div>
                  <Label className="text-xs text-muted-foreground">Activity Level</Label>
                  <div className="mt-2 flex flex-wrap gap-2">
                    {[
                      { label: "Any", value: null },
                      { label: "Low", value: 1 },
                      { label: "Steady", value: 2 },
                      { label: "Busy", value: 3 },
                      { label: "Very Busy", value: 4 },
                    ].map((option) => (
                      <button
                        key={option.label}
                        onClick={() => setActivityLevel(option.value)}
                        className={`rounded-lg border px-3 py-2 text-sm transition-colors ${
                          activityLevel === option.value
                            ? "border-accent bg-accent text-accent-foreground"
                            : "border-border bg-secondary text-muted-foreground hover:border-accent/50"
                        }`}
                      >
                        {option.label}
                      </button>
                    ))}
                  </div>
                </div>

                <div>
                  <Label className="text-xs text-muted-foreground">Moderation Strength</Label>
                  <div className="mt-2 flex flex-wrap gap-2">
                    {[
                      { label: "Any", value: null },
                      { label: "Light", value: 1 },
                      { label: "Standard", value: 2 },
                      { label: "Active", value: 3 },
                      { label: "Strict", value: 4 },
                    ].map((option) => (
                      <button
                        key={option.label}
                        onClick={() => setModerationLevel(option.value)}
                        className={`rounded-lg border px-3 py-2 text-sm transition-colors ${
                          moderationLevel === option.value
                            ? "border-accent bg-accent text-accent-foreground"
                            : "border-border bg-secondary text-muted-foreground hover:border-accent/50"
                        }`}
                      >
                        {option.label}
                      </button>
                    ))}
                  </div>
                </div>
              </div>
            )}
          </div>

          <Separator className="my-4" />

          <div className="mb-4">
            <button
              onClick={() => toggleSection("features")}
              className="flex w-full items-center justify-between py-2"
            >
              <span className="text-sm font-medium text-foreground">Features</span>
              {expandedSections.includes("features") ? (
                <ChevronUp className="h-4 w-4 text-muted-foreground" />
              ) : (
                <ChevronDown className="h-4 w-4 text-muted-foreground" />
              )}
            </button>
            {expandedSections.includes("features") && (
              <div className="mt-3 space-y-3">
                {features.map((feature) => (
                  <div key={feature.id} className="flex items-center gap-2">
                    <Checkbox
                      id={feature.id}
                      checked={selectedFeatures.includes(feature.id)}
                      onCheckedChange={() => toggleFeature(feature.id)}
                    />
                    <Label htmlFor={feature.id} className="cursor-pointer text-sm text-muted-foreground">
                      {feature.label}
                    </Label>
                  </div>
                ))}
              </div>
            )}
          </div>

          <div className="mt-6 flex gap-3">
            <Button
              variant="outline"
              className="flex-1 border-border text-foreground hover:bg-secondary"
              onClick={handleReset}
            >
              Reset
            </Button>
            <Button
              className="flex-1 bg-primary text-primary-foreground hover:bg-primary/90"
              onClick={handleApply}
            >
              Apply
            </Button>
          </div>
        </div>
      </aside>
    </>
  );
}
