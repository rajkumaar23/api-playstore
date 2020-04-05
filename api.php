<?php
/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

use Symfony\Component\DomCrawler\Crawler;

/**
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */
class API
{
    /**
     * @var string
     */
    private $playstoreURL;
    /**
     * @var Crawler
     */
    private $crawler;
    private $packageID;

    public function __construct($package)
    {
        $this->packageID = $package;
        $this->playstoreURL = 'https://play.google.com/store/apps/details?id=' . $this->packageID;
        if ($html = file_get_contents($this->playstoreURL)) {
            $this->crawler = new Crawler($html);
        } else {
            throw new Exception("Invalid Package ID");
        }
    }

    /**
     * @return mixed
     */
    public function getPackageID()
    {
        return $this->packageID;
    }

    /**
     * @return string
     */
    public function getVersion()
    {
        return $this->crawler->filter('.htlgb')->eq(6)->text();
    }

    /**
     * @return string
     */
    public function getInstalls()
    {
        return $this->crawler->filter('.htlgb')->eq(5)->text();
    }

    /**
     * @return string
     */
    public function getSize()
    {
        return $this->crawler->filter('.htlgb')->eq(3)->text();
    }

    /**
     * @return string
     */
    public function getLastUpdated()
    {
//        die(var_dump($this->crawler->filter('.EymY4b')));
        return $this->crawler->filter('.htlgb')->eq(1)->text();
    }

    /**
     * @return string
     */
    public function getRating()
    {
        return $this->crawler->filter('.BHMmbe')->eq(0)->text();
    }

    /**
     * @return string
     */
    public function getNoOfUsersRated()
    {
        return filter_var($this->crawler->filter('.EymY4b')->eq(0)->text(), FILTER_SANITIZE_NUMBER_INT);
    }

    /**
     * @return string
     */
    public function getDeveloper()
    {
        $size = sizeof($this->crawler->filter('.htlgb'));
        return $this->crawler->filter('.htlgb')->eq($size == 20 ? 17 : 18)->text();
    }
}
